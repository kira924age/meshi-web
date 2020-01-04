import React from 'react';
import { makeStyles } from '@material-ui/core/styles';
import Card from '@material-ui/core/Card';
import CardActionArea from '@material-ui/core/CardActionArea';
import CardActions from '@material-ui/core/CardActions';
import CardContent from '@material-ui/core/CardContent';
import Button from '@material-ui/core/Button';
import Typography from '@material-ui/core/Typography';
import Link from '@material-ui/core/Link';

const useStyles = makeStyles({
  card: {
    margin: 'auto',
    maxWidth: 345,
  },
  media: {
    height: 140,
  },
});

export default function MeshiCard(props) {
  const classes = useStyles();
  const req = 'https://www.google.com/search?q=調布+' + props.meshi;
  return (
    <Card className={classes.card}>
      <CardActionArea>
        <CardContent>
          <Typography gutterBottom variant="h5" component="h2">
            {props.meshi}
          </Typography>
          <Typography variant="body2" color="textSecondary" component="p">
          description of {props.meshi} 
          </Typography>
        </CardContent>
      </CardActionArea>
      <CardActions>
        <Button size="small" color="primary" style={{textTransform:"none"}}>
          <Link href= {req} target='_blank'>
          Search on Google
          </Link>
        </Button>
      </CardActions>
    </Card>
  );
}
