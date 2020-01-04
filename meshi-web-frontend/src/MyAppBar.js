import React from 'react';
import { makeStyles } from '@material-ui/core/styles';
import AppBar from '@material-ui/core/AppBar';
import Toolbar from '@material-ui/core/Toolbar';
import Typography from '@material-ui/core/Typography';
import GitHubIcon from '@material-ui/icons/GitHub';
import IconButton from '@material-ui/core/IconButton';
import Tooltip from '@material-ui/core/Tooltip';
import Link from '@material-ui/core/Link';

const useStyles = makeStyles(theme => ({
  root: {
    flexGrow: 1,
  },
  menuButton: {
    marginRight: theme.spacing(2),
  },
  title: {
    flexGrow: 1,
  },
  back: {
    backgroundColor: '#24292e',
  }
}));

export default function MyAppBar() {
  const classes = useStyles();

  return (
    <div className={classes.root}>
      <AppBar position="static" className={classes.back} >
        <Toolbar>
          <Typography variant="h5" className={classes.title}>
            meshi-web
          </Typography>
          <Tooltip title="GitHub Repository">
            <Link href="https://github.com/kira924age/meshi-web" target='_blank' color='inherit'>
              <IconButton edge="start" className={classes.menuButton} color="inherit" aria-label="menu">
              <GitHubIcon />
              </IconButton>
          </Link>
          </Tooltip>
        </Toolbar>
      </AppBar>
    </div>
  );
}
