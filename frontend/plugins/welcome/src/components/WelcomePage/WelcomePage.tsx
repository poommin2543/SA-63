import React from 'react';
import PropTypes from 'prop-types';
import AppBar from '@material-ui/core/AppBar';
import Avatar from '@material-ui/core/Avatar';
import Button from '@material-ui/core/Button';
import Grid from '@material-ui/core/Grid';
//import HelpIcon from '@material-ui/icons/Help';
import Hidden from '@material-ui/core/Hidden';
import IconButton from '@material-ui/core/IconButton';
import Link from '@material-ui/core/Link';
import MenuIcon from '@material-ui/icons/Menu';
import NotificationsIcon from '@material-ui/icons/Notifications';
import Tab from '@material-ui/core/Tab';
import Tabs from '@material-ui/core/Tabs';
import Toolbar from '@material-ui/core/Toolbar';
import Tooltip from '@material-ui/core/Tooltip';
import Typography from '@material-ui/core/Typography';
import { withStyles } from '@material-ui/core/styles';
import Paper from '@material-ui/core/Paper';
import TextField from '@material-ui/core/TextField';
//import Footer from './components/Footer';
//import Icon from '@material-ui/core/Icon';
import SaveIcon from '@material-ui/icons/Save';
import Autocomplete from '@material-ui/lab/Autocomplete';

const lightColor = 'rgba(255, 255, 255, 0.7)';

const top100Films = [
  { title: 'ใช้แล้วทิ้ง', year: 1994 },
  { title: 'นำกลับมาใช้ใหม่', year: 1972 }
];

const styles = (theme: { spacing: (arg0: number) => number; palette: { common: { white: any; }; }; }) => ({
  secondaryBar: {
    zIndex: 0,
  },
  menuButton: {
    marginLeft: -theme.spacing(1),
  },
  iconButtonAvatar: {
    padding: 4,
  },
  link: {
    textDecoration: 'none',
    color: lightColor,
    '&:hover': {
      color: theme.palette.common.white,
    },
  },
  button: {
    borderColor: lightColor,
  },
});

function Header(props: { classes: any; onDrawerToggle: any; }) {
  const { classes, onDrawerToggle } = props;

  return (
    <React.Fragment>
      <AppBar color="primary" position="sticky" elevation={0}>
        <Toolbar>
          <Grid container spacing={1} alignItems="center">
            <Hidden smUp>
              <Grid item>
                <IconButton
                  color="inherit"
                  aria-label="open drawer"
                  onClick={onDrawerToggle}
                  className={classes.menuButton}
                >
                  <MenuIcon />
                </IconButton>
              </Grid>
            </Hidden>
            <Grid item xs />
            <Grid item>
              <Link className={classes.link} href="#" variant="body2">
                log out
              </Link>
            </Grid>
            <Grid item>
              <Tooltip title="Alerts • No alerts">
                <IconButton color="inherit">
                  <NotificationsIcon />
                </IconButton>
              </Tooltip>
            </Grid>
            <Grid item>
              <IconButton color="inherit" className={classes.iconButtonAvatar}>
                <Avatar src="https://scontent.fnak3-1.fna.fbcdn.net/v/t1.0-1/cp0/c0.24.80.80a/p80x80/118251964_1523930951142174_1223119877856967161_o.jpg?_nc_cat=100&_nc_sid=7206a8&_nc_eui2=AeGud1VE_FvpszPeEGewB0_PjtLenCu9ryeO0t6cK72vJ0wzrk8_F0vaLzoBnmMEHlCtbaeFZ_vnHdYwBPs99lRu&_nc_ohc=woL9JSPw-psAX9bma7M&_nc_ht=scontent.fnak3-1.fna&oh=72514f0e3741918f167c24737161cb72&oe=5F7ADCB6" alt="Noom My Avatar" />
              </IconButton>
            </Grid>
          </Grid>
        </Toolbar>
      </AppBar>
      <AppBar
        component="div"
        className={classes.secondaryBar}
        color="primary"
        position="static"
        elevation={0}
      >
        <Toolbar>
          <Grid container alignItems="center" spacing={1}>
            <Grid item xs>
              <Typography color="inherit" variant="h1" component="h1">
                ระบบเครื่องมือแพทย์
              </Typography>
            </Grid>
            <Grid item>
              <Button className={classes.button} variant="outlined" color="inherit" size="small">
                นพ.ภูมิมินทร์ พินพิมาย
              </Button>
            </Grid>
            
          </Grid>
        </Toolbar>
      </AppBar>
      <AppBar
        component="div"
        className={classes.secondaryBar}
        color="primary"
        position="static"
        elevation={0}
      >
        <Tabs value={0} textColor="inherit">
          <Tab textColor="inherit" label="ADD Data" />
          <Tab textColor="inherit" label="DEL Data" />
          
          <Tab textColor="inherit" label="UPDATE Data" />
        </Tabs>
      </AppBar>
      
      <AppBar
        component="div"
        className={classes.secondaryBar}
        color="inherit"
        position="static"
        elevation={1}
      >
        <Toolbar>
          <Grid container alignItems="center" spacing={4}>
            <Grid item xs={12}></Grid>
            <Grid item xs={12}></Grid>
        
            
            <Grid item xs={2}></Grid>
            <Grid item xs={2}></Grid>
            <Grid item xs={2}>
              <Typography color="primary" variant="h6" component="h1">
                ID เครื่องมือแพทย์
              </Typography>
            </Grid>
            <Grid item xs={2}>
              <Paper className={classes.paper}>
                <TextField id="outlined-basic" label="กรุณาใส่ ID" variant="outlined"/></Paper>
            </Grid>
            <Grid item xs={2}></Grid>
            <Grid item xs={2}> </Grid>

            <Grid item xs={2}></Grid>
            <Grid item xs={2}></Grid>
            <Grid item xs={2}>
              <Typography color="primary" variant="h6" component="h1">
                ชื่อเครื่องมือแพทย์
              </Typography>
            </Grid>
            <Grid item xs={2}>
              <Paper className={classes.paper}>
                <TextField id="outlined-basic" label="กรุณาใส่ชื่อเครื่องมือ" variant="outlined"/></Paper>
            </Grid>
            <Grid item xs={2}></Grid>
            <Grid item xs={2}> </Grid>

            <Grid item xs={2}></Grid>
            <Grid item xs={2}></Grid>
            <Grid item xs={2}>
              <Typography color="primary" variant="h6" component="h1">
                ประเภทของเครื่องมือ
              </Typography>
            </Grid>
            <Grid item xs={2}>
            <Autocomplete
              id="combo-box-demo"
              options={top100Films}
              getOptionLabel={(option) => option.title}
              
              renderInput={(params) => <TextField {...params} label="โปรดระบุประเภท" variant="outlined" />}
            />
            </Grid>
            <Grid item xs={2}></Grid>
            <Grid item xs={2}> </Grid>

            <Grid item xs={2}></Grid>
            <Grid item xs={2}></Grid>
            <Grid item xs={2}>
              <Typography color="primary" variant="h6" component="h1">
                จำนวน/ชิ้น
              </Typography>
            </Grid>
            <Grid item xs={2}>
              <Paper className={classes.paper}>
                <TextField id="outlined-number" type='number' InputLabelProps={{
                  shrink: true,}}label="กรุณาใส่จำนวน" variant="outlined"/>
                  </Paper>
            </Grid>
            <Grid item xs={2}></Grid>
            <Grid item xs={2}> </Grid>


            <Grid item xs={2}></Grid>
            <Grid item xs={2}> </Grid>
            <Grid item xs={2}></Grid>
            <Grid item xs={2}>
              <Button
                variant="contained"
                color="primary"
                size="large"
                className={classes.button}
                startIcon={<SaveIcon />}
              >
                Save
              </Button>
            </Grid>
            <Grid item xs={2}></Grid>
            <Grid item xs={2}> </Grid>
            <Grid item xs={12}></Grid>
            <Grid item xs={12}></Grid>
            
    </Grid>
        </Toolbar>
      </AppBar>
      <AppBar
        component="div"
        className={classes.secondaryBar}
        color="primary"
        position="static"
        elevation={0}
      >
        <Toolbar>
          <Grid container alignItems="center" spacing={1}>
            <Grid item xs>
              
            </Grid>
            
            
          </Grid>
        </Toolbar>
      </AppBar>
           
    </React.Fragment>
  );
}

Header.propTypes = {
  classes: PropTypes.object.isRequired,
  onDrawerToggle: PropTypes.func.isRequired,
};

export default withStyles(styles)(Header);

