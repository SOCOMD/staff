import { h, Component } from 'preact'
import { Router } from 'preact-router'

import Header from './header'
import DashBoard from '../routes/Dashboard'
import Profile from '../routes/profile'
import Profile2 from '../routes/profile2'
import Login from '../routes/login'

export default class App extends Component<any, any> {

	currentUrl: string

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