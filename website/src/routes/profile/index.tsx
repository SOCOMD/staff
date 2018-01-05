import { h, Component, PreactHTMLAttributes } from 'preact';

import Dialog from 'preact-material-components/Dialog';
import 'preact-material-components/Dialog/style.css';

import List, { Divider } from 'preact-material-components/List';
import 'preact-material-components/Dialog/style.css';

import './style.css';


import { Grid, Paper, Button, Card, CardContent, CardActions, CardHeader, TextField, Typography, Select, MenuItem, Input, FormControl, FormHelperText, InputLabel } from "material-ui"
//import { DatePicker } from 'material-ui-pickers'
//GRPC Imports
import { grpc, BrowserHeaders, Code } from 'grpc-web-client'
import { staff } from '../../rpc/staff_pb_service'
import { User, GetUserRequest, UpdateUserRequest } from '../../rpc/staff_pb'

import { Router } from 'preact-router';
import { Event } from '_debugger';

export interface ProfileProps { profileID: string; self: boolean; }
export interface ProfileState { rpcUser: User; }

export default class Profile extends Component<ProfileProps, ProfileState> {


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
			debug: false,
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
				this.setState({ ...this.state, rpcUser: response })
				//this.setFieldValues(response)
			}
		})
	}

	updateUserData() {
		var request = new UpdateUserRequest;
		var token = sessionStorage.getItem("auth")
		request.setToken(token)
		request.setUser(this.state.rpcUser)
		console.log("User", this.state.rpcUser)
		grpc.unary(staff.UpdateUser, {
			debug: false,
			request: request,
			host: window.location.origin,
			onEnd: res => {
				const { status, statusMessage, headers, message, trailers } = res;
				if (status != Code.OK) {
					console.error(statusMessage);
				}
			}
		})
	}

	updateEmail = (e: any) => {
		var user = this.state.rpcUser
		//console.log("Email", e.target.value)
		user.setEmail(e.target.value)
		this.setState({ ...this.state, rpcUser: user })
	}

	handleGenderChange = e => {
		var user = this.state.rpcUser
		console.log("gender", e)
		user.setGender(e.target.value)
		this.setState({ ...this.state, rpcUser: user })
	}

	render() {
		var viewMode = false
		if (this.state.rpcUser == null) {
			return (<div></div>)
		}
		var user = this.state.rpcUser
		var self = this.props.self

		var spacing = { marginTop: '10px' }

		return (
			<div className="profile">
				<Card raised={true} style={{ padding: '20px' }}>
					<CardHeader title={<h1>Member - {user.getSteamid()}</h1>} />
					<CardContent>
						<Grid container direction="column">
							<h3>
								General
							</h3>
							<TextField fullwidth={true} helperTextPersistent={true} disabled={self == false} helperText="Email" value={user.getEmail()} onChange={this.updateEmail} />
							<TextField style={spacing} fullwidth={true} helperTextPersistent={true} disabled={self == false} helperText="Date of Birth" type="date" />

							{/*Cant work out why this wont work. Good luck chambers im getting on the piss
								https://material-ui-next.com/api/select/
								https://material-ui-next.com/demos/selects/						
								*/}
							<InputLabel htmlFor="gender">Gender</InputLabel>
							<Select
								value={user.getGender() == "" ? "Male" : user.getGender()}
								onChange={this.handleGenderChange}
								input={<Input id="gender" name="gender" />}
							>
								<MenuItem value="Male">Male</MenuItem>
								<MenuItem value="Female">Female</MenuItem>
								<MenuItem value="Other">Other</MenuItem>
							</Select>


							<span>Join Date: {user.getJoindate()}</span>
							<span>Active: {user.getActive() ? "YES" : "NO"}</span>
							<span>Admin: {user.getAdmin()}</span>

							<Divider />
							<h3>TeamSpeak</h3>
							<TextField style={spacing} fullwidth={true} helperTextPersistent={true} disabled={self == false} helperText="Teamspeak Unique ID" value={user.getTsuuid()} />
							<span style={spacing}>Name: {user.getTsname()}</span>
							<span>First Seen: {new Date(parseInt(user.getTscreated()) * 1000).toISOString().slice(0, 10)} </span>
							<span>Last  Seen: {new Date(parseInt(user.getTslastconnected()) * 1000).toISOString().slice(0, 10)}</span>
							<Divider />
							<h3>Ranks</h3>

						</Grid>
					</CardContent>
					<CardActions>
						<Button onClick={() => { this.scrollingDlg.MDComponent.show(); }}>Update</Button>
					</CardActions>
				</Card>

				<Dialog ref={scrollingDlg => { this.scrollingDlg = scrollingDlg; }}>
					<Dialog.Header>Updating Profile</Dialog.Header>
					<Dialog.Body>
						<p>Are you sure to want to update this information?</p>
					</Dialog.Body>
					<Dialog.Footer>
						<Dialog.FooterButton cancel={true}>Cancel</Dialog.FooterButton>
						<Dialog.FooterButton onClick={() => { this.updateUserData() }} accept={true}>Update</Dialog.FooterButton>
					</Dialog.Footer>
				</Dialog>
			</div >
		);
	}
}
