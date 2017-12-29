import { h, Component } from 'preact';
import Card from 'preact-material-components/Card';
import 'preact-material-components/Card/style.css';
import 'preact-material-components/Button/style.css';
import './style.css';

import Icon from 'preact-material-components/Icon';
import 'preact-material-components/Icon/style.css';

export default class Home extends Component<any, any> {
	render() {
		return (
			<div className="home">
				<h1>Home route</h1>
				<Icon>menu</Icon>
				<Card>
					<Card.Primary>
						<Card.Title>Home card</Card.Title>
						<Card.Subtitle>Welcome to home route</Card.Subtitle>
					</Card.Primary>
					<Card.SupportingText>
						Sed ut perspiciatis unde omnis iste natus error sit voluptatem accusantium doloremque laudantium, totam rem aperiam, eaque ipsa quae ab illo inventore veritatis et quasi architecto beatae vitae dicta sunt explicabo. Nemo enim ipsam voluptatem quia voluptas sit aspernatur aut odit aut fugit, sed quia consequuntur magni dolores eos qui ratione voluptatem sequi nesciunt.
					</Card.SupportingText>
					<Card.Actions>
						<Card.Action>OKAY</Card.Action>
					</Card.Actions>
				</Card>
			</div>
		);
	}
}