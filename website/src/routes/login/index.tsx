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
				Login Screen
			</div>
		)
	}
}