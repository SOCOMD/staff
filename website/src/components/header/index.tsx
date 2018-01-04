import { h, Component } from 'preact';
import { Link } from 'preact-router'

import PropTypes from 'prop-types';
import { AppBar, Toolbar, Typography, Button, Divider } from 'material-ui'
const styles = {
	root: {
		width: '100%',
	},
};

export default class Header extends Component<any, any> {
	render() {

		return (
			<div style={styles} >
				<AppBar position="static" color="primary">
					<Toolbar>
						<Typography type="title" color="inherit">
							Socomd
						</Typography>
						<Link activeClassName="active" href="/">Home</Link>
						<Link activeClassName="active" href="/profile">Profile</Link>
					</Toolbar>
				</AppBar>
			</div>
		);
	}
}
