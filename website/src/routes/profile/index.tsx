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
import { members } from '../../rpc/staff_pb_service'
import { User, GetUserMessage } from '../../rpc/staff_pb'

import { Router } from 'preact-router';

export default class Profile extends Component<any, any> {

	timer: any

	scrollingDlg: any
	state = {
		time: Date.now(),
		count:0,
		rpcUser:null
	};

	// gets called when this route is navigated to
	componentDidMount() {
		// start a timer for the clock:
		this.timer = setInterval(this.updateTime, 1000);
		this.requestUserData();
	}

	// gets called just before navigating away from the route
	componentWillUnmount() {
		clearInterval(this.timer);
	}

	// update the current time
	updateTime = () => {
		this.setState({ time: Date.now() });
	};

	increment = () => {
		this.setState({ count: this.state.count + 1 });
	};

	updateUser = (usr : User) => {
		this.setState({rpcUser: usr})
	};

	extractProfileID(){
		var url = window.location.href
		var lastIdx = url.lastIndexOf('/')
		var profileId = url.substr(lastIdx + 1, url.length - lastIdx)
		return profileId
	}

	requestUserData() {

		//TODO: GET USER ID FROM ROUTE
		var request = new GetUserMessage;
		var profileID = this.extractProfileID()
		request.setId(profileID);

		grpc.unary(members.GetUser, {
			debug:true,
			request: request,
			host: process.env.webgrpc_host,
			onEnd: res => {
				const { status, statusMessage, headers, message, trailers } = res;
					if(status != Code.OK || !message) {
						console.log(statusMessage);
					}

					var response = message as User
					if(response == null) {
						return						
					}

					this.setState({rpcUser: response.toObject()})
				}
			}
		)
	}

	formatDate(date) {
		var d = new Date(date),
			month = '' + (d.getMonth() + 1),
			day = '' + d.getDate(),
			year = d.getFullYear();
	
		if (month.length < 2) month = '0' + month;
		if (day.length < 2) day = '0' + day;
	
		return [year, month, day].join('-');
	}

	renderProfileCard({rpcUser}) {

		var isDisabled = true
		if(rpcUser == null) {
			return;
		}

		console.log()
		
		return (
			<div>
				<Card>
					<Card.Primary>
						<h1>Profile:{rpcUser.tsname}</h1>
						<LayoutGrid>
							<LayoutGrid.Inner>
							<LayoutGrid.Cell cols={4}>
									<TextField fullwidth={true} helperTextPersistent={true} disabled={isDisabled} helperText="ID" value={rpcUser.id}/>
								</LayoutGrid.Cell>
								<LayoutGrid.Cell cols={4}>
									<TextField fullwidth={true} helperTextPersistent={true} disabled={isDisabled} helperText="Teamspeak ID" value={rpcUser.tsdbid}/>
								</LayoutGrid.Cell>
								<LayoutGrid.Cell cols={4}>
									<TextField fullwidth={true} helperTextPersistent={true} disabled={isDisabled} helperText="Teamspeak Unique ID" value={rpcUser.tsuuid}/>
								</LayoutGrid.Cell>
								<LayoutGrid.Cell cols={4}>
									<TextField fullwidth={true} helperTextPersistent={true} disabled={isDisabled} helperText="Teamspeak Name" value={rpcUser.tsname}/>
								</LayoutGrid.Cell>
								<LayoutGrid.Cell cols={4}>
									<TextField fullwidth={true} helperTextPersistent={true} disabled={isDisabled} helperText="Teamspeak Created" value={new Date(parseInt(rpcUser.tscreated) * 1000).toISOString().slice(0, 10)} type="date"/>
								</LayoutGrid.Cell>
								<LayoutGrid.Cell cols={4}>
									<TextField fullwidth={true} helperTextPersistent={true} disabled={isDisabled} helperText="Teamspeak Last Connected" value={new Date(parseInt(rpcUser.tslastconnected) * 1000).toISOString().slice(0, 10)} type="date"/>
								</LayoutGrid.Cell>
								<LayoutGrid.Cell cols={4}>
									<TextField fullwidth={true} helperTextPersistent={true} disabled={isDisabled} helperText="Email" value={rpcUser.email}/>
								</LayoutGrid.Cell>
								<LayoutGrid.Cell cols={4}>
									<TextField fullwidth={true} helperTextPersistent={true} disabled={isDisabled} helperText="Join Date" value={rpcUser.joindate} type="date"/>
								</LayoutGrid.Cell>
								<LayoutGrid.Cell cols={4}>
									<TextField fullwidth={true} helperTextPersistent={true} disabled={isDisabled} helperText="Date of Birth" value={rpcUser.dob} type="date"/>
								</LayoutGrid.Cell>
								<LayoutGrid.Cell cols={4}>
									<TextField fullwidth={true} helperTextPersistent={true} disabled={isDisabled} helperText="Gender" value={rpcUser.gender}/>
								</LayoutGrid.Cell>
								<LayoutGrid.Cell cols={4}>
									<TextField fullwidth={true} helperTextPersistent={true} disabled={isDisabled} helperText="Active" value={rpcUser.active}/>
								</LayoutGrid.Cell>
								<LayoutGrid.Cell cols={4}>
									<TextField fullwidth={true} helperTextPersistent={true} disabled={isDisabled} helperText="Admin" value={rpcUser.admin} type="number"/>
								</LayoutGrid.Cell>
							</LayoutGrid.Inner>
						</LayoutGrid>
					</Card.Primary>
				</Card>

				<Dialog ref={scrollingDlg=>{this.scrollingDlg=scrollingDlg;}}>
					<Dialog.Header>Updating Profile</Dialog.Header>
					<Dialog.Body>
						<p>Are you sure to want to update this information?</p>
					</Dialog.Body>
					<Dialog.Footer>
						<Dialog.FooterButton cancel={true}>Cancel</Dialog.FooterButton>
						<Dialog.FooterButton accept={true}>Update</Dialog.FooterButton>
					</Dialog.Footer>
				</Dialog>
			</div>
		)
	}

	// Note: `user` comes from the URL, courtesy of our router
	render({ user }, {time, count, rpcUser}) {
		return (
			<div className="profile">
				<this.renderProfileCard rpcUser={rpcUser}/>
			</div>
		);
	}
}
