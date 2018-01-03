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
import { User, GetUserMessage } from '../../rpc/staff_pb'

import { Router } from 'preact-router';

export default class Profile extends Component<any, any> {

	timer: any

	scrollingDlg: any
	state = {
		rpcUser: null
	};

	// gets called when this route is navigated to
	componentDidMount() {
		this.requestUserData();
	}

	// gets called just before navigating away from the route
	componentWillUnmount() {
		clearInterval(this.timer);
	}

	updateUser = (usr: User) => {
		this.setState({ rpcUser: usr })
	};

	extractProfileID() {
		var url = window.location.href
		var lastIdx = url.lastIndexOf('/')
		var profileId = url.substr(lastIdx + 1, url.length - lastIdx)
		return profileId
	}

	requestUserData() {

		var request = new GetUserMessage;
		var profileID = this.extractProfileID()
		request.setId(profileID);

		grpc.unary(staff.GetUser, {
			debug: true,
			request: request,
			host: window.location.origin,
			onEnd: res => {
				const { status, statusMessage, headers, message, trailers } = res;
				if (status != Code.OK || !message) {
					console.log(statusMessage);
				}

				var response = message as User
				if (response == null) {
					return
				}

				this.setFieldValues(response)
			}
		}
		)
	}

	setFieldValues(rpcUser: User) {
		if (rpcUser == null) {
			return;
		}

		this.setState({ rpcUser: rpcUser.toObject() })

		document.getElementById("tsuuid").setAttribute("value", this.state.rpcUser.tsuuid)
		document.getElementById("tsname").setAttribute("value", this.state.rpcUser.tsname)
		document.getElementById("tscreated").setAttribute("value", new Date(parseInt(this.state.rpcUser.tscreated) * 1000).toISOString().slice(0, 10))
		document.getElementById("tslastconnected").setAttribute("value", new Date(parseInt(this.state.rpcUser.tslastconnected) * 1000).toISOString().slice(0, 10))
		document.getElementById("email").setAttribute("value", this.state.rpcUser.email)
		document.getElementById("joindate").setAttribute("value", this.state.rpcUser.joindate)
		document.getElementById("dob").setAttribute("value", this.state.rpcUser.dob)
		document.getElementById("gender").setAttribute("value", this.state.rpcUser.gender)
		document.getElementById("active").setAttribute("value", this.state.rpcUser.active)
		document.getElementById("admin").setAttribute("value", this.state.rpcUser.admin)

		//document.getElementById("performUpdate").addEventListener("click", (e:Event) => this.onFormSubmit());
	}

	renderProfileCard({ user }) {

		var viewMode = true

		return (
			<div>
				<Card>
					<Card.Primary>
						<h1>Member - {user}</h1>
						<LayoutGrid>
							<LayoutGrid.Inner>
								<LayoutGrid.Cell cols={4}>
									<TextField id="tsname" fullwidth={true} helperTextPersistent={true} disabled={true} helperText="Teamspeak Name" />
								</LayoutGrid.Cell>
								<LayoutGrid.Cell cols={4}>
									<TextField id="email" fullwidth={true} helperTextPersistent={true} disabled={viewMode} helperText="Email" />
								</LayoutGrid.Cell>
								<LayoutGrid.Cell cols={4}>
									<TextField id="tsuuid" fullwidth={true} helperTextPersistent={true} disabled={viewMode} helperText="Teamspeak Unique ID" />
								</LayoutGrid.Cell>
								<LayoutGrid.Cell cols={4}>
									<TextField id="tscreated" fullwidth={true} helperTextPersistent={true} disabled={true} helperText="Teamspeak Created" type="date" />
								</LayoutGrid.Cell>
								<LayoutGrid.Cell cols={4}>
									<TextField id="tslastconnected" fullwidth={true} helperTextPersistent={true} disabled={true} helperText="Teamspeak Last Connected" type="date" />
								</LayoutGrid.Cell>
								<LayoutGrid.Cell cols={4}>
									<TextField id="joindate" fullwidth={true} helperTextPersistent={true} disabled={true} helperText="Join Date" type="date" />
								</LayoutGrid.Cell>
								<LayoutGrid.Cell cols={4}>
									<TextField id="dob" fullwidth={true} helperTextPersistent={true} disabled={true} helperText="Date of Birth" type="date" />
								</LayoutGrid.Cell>
								<LayoutGrid.Cell cols={4}>
									<TextField id="gender" fullwidth={true} helperTextPersistent={true} disabled={viewMode} helperText="Gender" />
								</LayoutGrid.Cell>
								<LayoutGrid.Cell cols={4}>
									<TextField id="active" fullwidth={true} helperTextPersistent={true} disabled={viewMode} helperText="Active" />
								</LayoutGrid.Cell>
								<LayoutGrid.Cell cols={4}>
									<TextField id="admin" fullwidth={true} helperTextPersistent={true} disabled={viewMode} helperText="Admin" type="number" />
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
		)
	}

	// Note: `user` comes from the URL, courtesy of our router
	render({ user }) {
		return (
			<div className="profile">
				<this.renderProfileCard user={user} />
			</div>
		);
	}
}
