import { h, Component } from 'preact';
import { Link } from 'preact-router'

import PropTypes from 'prop-types';
import { AppBar, Toolbar, Typography, Button, Divider } from 'material-ui'
const styles = {
	root: {
		width: '100%',
	},
};

export interface HeaderProps { admin: Number }

export default class Header extends Component<HeaderProps, any> {

	render() {

		return (
			<div style={styles} >
				<AppBar position="static" color="primary">
					<Toolbar>
						<Typography type="title" color="inherit">
							Socomd
						</Typography>
						<Link activeClassName="active" href={this.props.admin ? "/" : "/profile"}>Home</Link>
						{this.props.admin ? <Link activeClassName="active" href="/profile">Profile</Link> : null}
					</Toolbar>
				</AppBar>
			</div>
		);
	}
}
