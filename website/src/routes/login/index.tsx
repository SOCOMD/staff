import { h, Component } from 'preact';
import { Button } from 'material-ui'
export default class Login extends Component<any, any> {
	username: string
	password: string

	login = () => {
		// send http request and find out. 
	}

	render() {
		return (
			<div>
				<div>
					Login Screen
				</div>
				<a href="/steamlogin"><Button raised color="primary">Steam Login (insert steam logo)</Button></a>
			</div>
		)
	}
}