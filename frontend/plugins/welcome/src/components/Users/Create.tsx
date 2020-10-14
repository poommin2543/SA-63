import React, { useState,Component } from 'react';
import { createStyles, makeStyles, Theme } from '@material-ui/core/styles';
import AppBar from '@material-ui/core/AppBar';
import Toolbar from '@material-ui/core/Toolbar';
import Typography from '@material-ui/core/Typography';
import IconButton from '@material-ui/core/IconButton';
import MenuIcon from '@material-ui/icons/Menu';
import AccountCircle from '@material-ui/icons/AccountCircle';
import Grid from '@material-ui/core/Grid';
import MenuItem from '@material-ui/core/MenuItem';
import Menu from '@material-ui/core/Menu';
import Button from '@material-ui/core/Button';
import AccountBoxIcon from '@material-ui/icons/AccountBox';
import TextField from '@material-ui/core/TextField';
import Autocomplete from '@material-ui/lab/Autocomplete';
import TableCell from '@material-ui/core/TableCell';
import FormControl from '@material-ui/core/FormControl';

import Hidden from '@material-ui/core/Hidden';
import DeleteIcon from '@material-ui/icons/Delete';
import { Link as RouterLink } from 'react-router-dom';
import Link from '@material-ui/core/Link';
import Tooltip from '@material-ui/core/Tooltip';
import Avatar from '@material-ui/core/Avatar';
import Tab from '@material-ui/core/Tab';
import Tabs from '@material-ui/core/Tabs';
import SvgIcon from '@material-ui/core/SvgIcon';
import Paper from '@material-ui/core/Paper';
import SaveIcon from '@material-ui/icons/Save';
import { DefaultApi } from '../../api/apis';




const lightColor = 'rgba(255, 255, 255, 0.7)';

const useStyles = makeStyles((theme: Theme) =>
  createStyles({
    root: {
      flexGrow: 1,
    },
    menuButton: {
      marginRight: theme.spacing(2),
    },
    iconButtonAvatar: {
      padding: 4,
    },
    secondaryBar: {
      zIndex: 0,
    },
    button: {
      borderColor: lightColor,
    },
    title: {
      flexGrow: 1,
    },
  }),
);

const name = [
  { title: 'นพ. ภูมิมินทร์ ' },
  { title: 'นพ. พินพิมาย' }
];

const top100Films = [
  { title: 'ใช้แล้วทิ้ง' },
  { title: 'นำกลับมาใช้ใหม่' }
];

const initialSystemequipmentState = {
  //medicalequipment: 'noom',
  addedTime : '',
 };


export default function MenuAppBar() {
  
  const [anchorEl, setAnchorEl] = React.useState<null | HTMLElement>(null);
  const open = Boolean(anchorEl);

  
  const classes = useStyles();
  const profile = { givenName: 'to Software Analysis 63' };
  const api = new DefaultApi();
  
  const [systemequipment, setSystemequipment] = useState(initialSystemequipmentState);
  const [status, setStatus] = useState(false);
  const [alert, setAlert] = useState(true);
  
  const handleInputChange = (event: any) => {
    const { id, value } = event.target;
    setSystemequipment({ ...systemequipment, [id]: value });
  };
  
  const CreateSystemequipment = async () => {
    const res:any = await api.createSystemequipment({ systemequipment });
    setStatus(true);
    if (res.id !== ''){
      setAlert(true);
    } else {
      setAlert(false);
    }
    const timer = setTimeout(() => {
      setStatus(false);
    }, 1000);
  };


function HomeIcon(props:any) {
    return (
      <SvgIcon {...props}>
        <path d="M10 20v-6h4v6h5v-8h3L12 3 2 12h3v8z" />
      </SvgIcon>
    );
  }

  return (
    <div className={classes.root}>
      <AppBar color="primary" position="sticky" elevation={0}>
        <Toolbar>
          <Grid container spacing={1} alignItems="center">
            <Hidden smUp>
            <Grid item>
                
              </Grid>
            </Hidden>
            <Grid item xs />
            <Grid item>
            <Link
            component="button"
            variant="body2"
            onClick={() => {
              console.info("I'm a button.");
            }}
          >
            Logout
            </Link>
            </Grid>
            <Grid item>
              
            </Grid>
            <Grid item>
              <IconButton color="inherit" className={classes.iconButtonAvatar}>
                <Avatar src="https://shorturl.at/ixKX1" alt="NoomAvatar" />
              </IconButton>
            </Grid>
          </Grid>
        </Toolbar>
      </AppBar>

      <AppBar
        component="div"
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
                <IconButton 
                style={{ marginLeft: 20 }}
                component={RouterLink}
                to="/"
                >    
                <HomeIcon color="inherit" />
                </IconButton>
                </Grid>  
            
            <Grid item>
            <Autocomplete
              id="combo-box-demo"
              options={name}
              getOptionLabel={(option) => option.title}
              style={{ width: 200 }}
              
              renderInput={(params) => <TextField {...params} label=""  variant="outlined" />}
                  
            />
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
                ชื่อเครื่องมือแพทย์
              </Typography>
            </Grid>
            <Grid item xs={2}>
              <Paper >
              <TextField
               id="addedTime"
               label="ชื่อเครื่องมือแพทย์"
               variant="outlined"
               type="string"
               size="medium"
               value={systemequipment.addedTime}
               onChange={handleInputChange}
             /></Paper>
            </Grid>
            <Grid item xs={2}></Grid>
            <Grid item xs={2}> </Grid>

            <Grid item xs={2}></Grid>
            <Grid item xs={2}></Grid>
            <Grid item xs={2}>
              <Typography color="primary" variant="h6" component="h1">
                ประเภทของอุปกรณ์
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
              <Paper >
                <TextField id="outlined-number" type='number' InputLabelProps={{
                  shrink: true,}}label="กรุณาใส่จำนวน" variant="outlined"
                  />
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
                onClick={() => {
                  CreateSystemequipment();
                }}
                startIcon={<SaveIcon 
                />}
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
            new Date() 
            </Grid>
          </Grid>
        </Toolbar>
      </AppBar>
    </div>
  );
}