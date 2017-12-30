import { h, Component } from 'preact';
import Card from 'preact-material-components/Card';
import 'preact-material-components/Card/style.css';
import 'preact-material-components/Button/style.css';
import './style.css';

import Icon from 'preact-material-components/Icon';
import 'preact-material-components/Icon/style.css';

export default class DashBoard extends Component<any, any> {
	render() {
		return (
			<div className="home">
				<h1>Dashboard</h1>
				<Card>
					<Card.Primary>
						<Card.Title>Dashboard</Card.Title>
						<Card.Subtitle>Welcome to the dashboard</Card.Subtitle>
					</Card.Primary>
					<Card.SupportingText>

					</Card.SupportingText>
					<Card.Actions>
						<Card.Action>OKAY</Card.Action>
					</Card.Actions>
				</Card>
			</div>
		);
	}
}