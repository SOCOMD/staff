import { h, Component } from 'preact';

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
				<a href="/steamlogin"><button>Steam Login (insert steam logo)</button></a>
			</div>
		)
	}
}