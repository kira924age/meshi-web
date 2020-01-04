import React from 'react';
import Typography from '@material-ui/core/Typography';
import Button from '@material-ui/core/Button';
import PropTypes from 'prop-types';
import { withStyles } from '@material-ui/styles';
import axios from 'axios';
import MeshiCard from './MeshiCard'

const styles = theme => ({
  root: {
    margin: '20px',
    flexGrow: 1,
  },
  menuButton: {
    margin: '10px',
    backgroundColor: '#24292e',
    color: '#fff',
    textTransform: 'none',
  },
  title: {
    flexGrow: 1,
  },
  center: {
    textAlign: 'center',
  },
});

class Contents extends React.Component {
  constructor(props) {
    super(props);
    this.state = {
      meshi: 'shokujin',
    };
    // This binding is necessary to make `this` work in the callback
    this.handleGetMeshi = this.handleGetMeshi.bind(this);
  }
  handleGetMeshi() {
    axios
      .get('http://kira924age.com:8080/meshi')
      .then((results) => {
        const data = results.data;
        this.setState(state => ({
          meshi: data
        }));
      },
      )
      .catch(() => {
        console.log('通信に失敗しました。');
      });
  }

  render() {
    const { classes } = this.props;
    return (
      <div className={classes.root}>
        <Typography variant="h4" align='center'>
        decide meshi
        </Typography>
        <div className={classes.center}>
          <Button variant="contained" size='large' className={classes.menuButton} onClick={() => this.handleGetMeshi()}>
          <Typography>
            meshiコマンドを叩く
          </Typography>
          </Button>
        </div>
        <MeshiCard meshi= {this.state.meshi} />
      </div>
    );
  }
}

Contents.propTypes = {
  classes: PropTypes.object.isRequired,
};

export default withStyles(styles)(Contents);

