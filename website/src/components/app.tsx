import { h, Component } from 'preact'
import { Router, route } from 'preact-router'

import Header from './header'
import DashBoard from '../routes/Dashboard'
import Profile from '../routes/profile'
import Profile2 from '../routes/profile2'
import Login from '../routes/login'

import { grpc, BrowserHeaders, Code } from 'grpc-web-client'
import { staff } from '../rpc/staff_pb_service'
import { GetAuthStatusRequest, GetAuthStatusResult } from '../rpc/staff_pb'

export default class App extends Component<any, any> {

	currentUrl: string

	componentDidMount() {
		console.log("Url:", this.currentUrl)
		if (this.currentUrl != "/login") {
			var token = sessionStorage.getItem("auth")
			console.log("Token", token)
			if (token == null || token.length == 0) {
				console.log("Invalid token")
				route("/login", true)
			}
			var request = new GetAuthStatusRequest
			request.setToken(token)

			grpc.unary(staff.AuthStatus, {
				debug: true,
				request: request,
				host: window.location.origin,
				onEnd: res => {
					const { status, statusMessage, headers, message, trailers } = res;
					if (status != Code.OK || !message) {
						route("/login", true)
						//console.error(statusMessage);
					}

					var response = message as GetAuthStatusResult
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
					//this.setState({ ...this.state, rpcUser: response.toObject() })
					//this.setFieldValues(response)
				}
			})
		}
	}
	handleRoute = (e: any) => {
		this.currentUrl = e.url
	}
	render() {
		return (
			<div>
				<Header />
				<Router onChange={this.handleRoute}>
					<Login path="/login" />
					<DashBoard path="/" />
					<Profile2 path="/profile/" profileID="" />
					<Profile2 path="/profile/:profileID" profileID="" />
				</Router>
			</div>
		)
	}
}