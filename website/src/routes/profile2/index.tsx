import { h, Component } from 'preact';

import Button from 'preact-material-components/Button';
import 'preact-material-components/Button/style.css';

import Card from 'preact-material-components/Card';
import 'preact-material-components/Card/style.css';

import TextField from 'preact-material-components/TextField';
import 'preact-material-components/TextField/style.css';

import LayoutGrid from 'preact-material-components/LayoutGrid';
import 'preact-material-components/LayoutGrid/style.css';

import Dialog from 'preact-material-components/Dialog';
import 'preact-material-components/Dialog/style.css';

import List from 'preact-material-components/List';
import 'preact-material-components/Dialog/style.css';

import './style.css';

//GRPC Imports
import { grpc, BrowserHeaders, Code } from 'grpc-web-client'
import { staff } from '../../rpc/staff_pb_service'
import { User, GetUserRequest } from '../../rpc/staff_pb'

import { Router } from 'preact-router';

export interface ProfileProps { profileID: string; }
export interface ProfileState { rpcUser: User.AsObject; }

export default class Profile2 extends Component<ProfileProps, ProfileState> {


	scrollingDlg: any

	// gets called when this route is navigated to
	componentDidMount() {
		this.requestUserData();
	}

	// gets called just before navigating away from the route
	componentWillUnmount() {

	}

	requestUserData() {

		var request = new GetUserRequest;

		var token = sessionStorage.getItem("auth")
		request.setToken(token)
		console.log("profileid", this.props.profileID)
		request.setSearch(this.props.profileID);
		if (this.props.profileID.length == 0) {
			request.setType(GetUserRequest.searchType.TOKEN)
		} else {
			request.setType(GetUserRequest.searchType.STEAMID)
		}
		grpc.unary(staff.GetUser, {
			debug: true,
			request: request,
			host: window.location.origin,
			onEnd: res => {
				const { status, statusMessage, headers, message, trailers } = res;
				if (status != Code.OK || !message) {
					console.error(statusMessage);
				}

				var response = message as User
				if (response == null) {
					return
				}
				/*
					...this.state - spreads all values inside of it out as if they were
					all manually entered. the following ',rpcUser:' overrides the previous
					rpcUser that ...this.state put. This way we can keep all over state
					values presuming we have others, and overwrite only the thing we want
					to change. 
				*/
				this.setState({ ...this.state, rpcUser: response.toObject() })
				//this.setFieldValues(response)
			}
		})
	}


	render() {
		var viewMode = false
		if (this.state.rpcUser == null) {
			return (<div></div>)
		}
		var user = this.state.rpcUser
		console.log("rpcUser", this.state.rpcUser)
		return (
			<div className="profile">
				<Card>
					<Card.Primary>
						<h1>Member - {user.steamid}</h1>
						<LayoutGrid>
							<LayoutGrid.Inner>
								<LayoutGrid.Cell cols={4}>
									<TextField
										id="tsname"
										fullwidth={true}
										helperTextPersistent={true}
										disabled={true}
										helperText="Teamspeak Name"
										value={user.tsname}
									/>
								</LayoutGrid.Cell>
								<LayoutGrid.Cell cols={4}>
									<TextField
										id="email"
										fullwidth={true}
										helperTextPersistent={true}
										disabled={viewMode}
										helperText="Email"
										value={user.email}
									/>
								</LayoutGrid.Cell>
								<LayoutGrid.Cell cols={4}>
									<TextField
										id="tsuuid"
										fullwidth={true}
										helperTextPersistent={true}
										disabled={viewMode}
										helperText="Teamspeak Unique ID"
										value={user.tsuuid}
									/>
								</LayoutGrid.Cell>
								<LayoutGrid.Cell cols={4}>
									<TextField
										id="tscreated"
										fullwidth={true}
										helperTextPersistent={true}
										disabled={true}
										helperText="Teamspeak Created"
										type="date"
										value={new Date(parseInt(user.tscreated) * 1000).toISOString().slice(0, 10)}
									/>
								</LayoutGrid.Cell>
								<LayoutGrid.Cell cols={4}>
									<TextField
										id="tslastconnected"
										fullwidth={true}
										helperTextPersistent={true}
										disabled={true}
										helperText="Teamspeak Last Connected"
										type="date"
										value={new Date(parseInt(user.tslastconnected) * 1000).toISOString().slice(0, 10)}
									/>
								</LayoutGrid.Cell>
								<LayoutGrid.Cell cols={4}>
									<TextField
										id="joindate"
										fullwidth={true}
										helperTextPersistent={true}
										disabled={true}
										helperText="Join Date"
										type="date"
										value={user.joindate}
									/>
								</LayoutGrid.Cell>
								<LayoutGrid.Cell cols={4}>
									<TextField
										id="dob"
										fullwidth={true}
										helperTextPersistent={true}
										disabled={true}
										helperText="Date of Birth"
										value={user.dob}
									/>
								</LayoutGrid.Cell>
								<LayoutGrid.Cell cols={4}>
									<TextField
										id="gender"
										fullwidth={true}
										helperTextPersistent={true}
										disabled={viewMode}
										helperText="Gender"
										value={user.gender}
									/>
								</LayoutGrid.Cell>
								<LayoutGrid.Cell cols={4}>
									<TextField
										id="active"
										fullwidth={true}
										helperTextPersistent={true}
										disabled={viewMode}
										helperText="Active"
										value={user.active ? "Yes" : "No"}
									/>
								</LayoutGrid.Cell>
								<LayoutGrid.Cell cols={4}>
									<TextField
										id="admin"
										fullwidth={true}
										helperTextPersistent={true}
										disabled={viewMode}
										helperText="Admin"
										type="number"
										value={user.admin.toString()}
									/>
								</LayoutGrid.Cell>
							</LayoutGrid.Inner>
						</LayoutGrid>
					</Card.Primary>
					<Card.Action onClick={() => { this.scrollingDlg.MDComponent.show(); }}>Update</Card.Action>
				</Card>

				<Dialog ref={scrollingDlg => { this.scrollingDlg = scrollingDlg; }}>
					<Dialog.Header>Updating Profile</Dialog.Header>
					<Dialog.Body>
						<p>Are you sure to want to update this information?</p>
					</Dialog.Body>
					<Dialog.Footer>
						<Dialog.FooterButton cancel={true}>Cancel</Dialog.FooterButton>
						<Dialog.FooterButton id="performUpdate" accept={true}>Update</Dialog.FooterButton>
					</Dialog.Footer>
				</Dialog>
			</div>
		);
	}
}
