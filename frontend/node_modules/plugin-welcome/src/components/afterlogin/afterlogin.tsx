import React, { } from 'react';
import { Link as RouterLink } from 'react-router-dom';
import ComponanceTable from '../Table';
import Button from '@material-ui/core/Button';
import { createStyles, makeStyles, Theme } from '@material-ui/core/styles';
import AppBar from '@material-ui/core/AppBar';
import Toolbar from '@material-ui/core/Toolbar';
import Typography from '@material-ui/core/Typography';
import Grid from '@material-ui/core/Grid';


import { Hidden } from '@material-ui/core';
 



const useStyles = makeStyles((theme: Theme) =>
  createStyles({
    root: {
      flexGrow: 1,
    },
    menuButton: {
      marginRight: theme.spacing(2),
    },
    title: {
      flexGrow: 1,
      
    },
    
  }),
);



export default function MenuAppBar() {
  const classes = useStyles();

  return (
    <div className={classes.root}>
      
      <AppBar  color="secondary" position="sticky" elevation={0}>
        <Toolbar>
          <Grid container spacing={1} alignItems="center">
            <Hidden smUp>
            <Grid item>
                
              </Grid>
            </Hidden>
            <Grid item xs />
            <Grid item>
            
            </Grid>
            
            <Grid item>
            <Button  variant="outlined" color="inherit" 
            size="small" component={RouterLink}
            to="/adddata">
                เพิ่มข้อมูล
              </Button>
                </Grid>
                <Grid item>
            <Button variant="outlined" color="inherit" 
            size="small" component={RouterLink}
            to="/">
                logout
              </Button>
                </Grid>  
            <Grid item>
            
            </Grid>
            
          </Grid>
        </Toolbar>
      </AppBar>
      <AppBar
        component="div"
        color="secondary"
        position="static"
        elevation={0}
      >
        <Toolbar>
          <Grid container alignItems="center" spacing={1}>
            <Grid item xs>
              <Typography color="inherit" variant="h2" component="h2">
                รายการบันทึกข้อมูล
              </Typography>
            </Grid>
            <Grid item>
            
                </Grid>  
            
            <Grid item>
            
            </Grid>
            
          </Grid>
        </Toolbar>
      </AppBar>
        <ComponanceTable></ComponanceTable>
      
    </div>
  );
}
 
