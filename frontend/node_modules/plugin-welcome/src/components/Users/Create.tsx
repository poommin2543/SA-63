import React, { useState,Component, useEffect } from 'react';
import { createStyles, makeStyles, Theme } from '@material-ui/core/styles';
import AppBar from '@material-ui/core/AppBar';
import Toolbar from '@material-ui/core/Toolbar';
import Typography from '@material-ui/core/Typography';
import IconButton from '@material-ui/core/IconButton';
import Grid from '@material-ui/core/Grid';
import MenuItem from '@material-ui/core/MenuItem';
import Button from '@material-ui/core/Button';
import Hidden from '@material-ui/core/Hidden';
import { Link as RouterLink } from 'react-router-dom';
import Link from '@material-ui/core/Link';
import Avatar from '@material-ui/core/Avatar';
import Tab from '@material-ui/core/Tab';
import Tabs from '@material-ui/core/Tabs';
import SvgIcon from '@material-ui/core/SvgIcon';
import SaveIcon from '@material-ui/icons/Save';
import { DefaultApi } from '../../api/apis';
import { EntPhysician } from '../../api/models/EntPhysician';
import ComponanceTable from '../Table';   
import Select from '@material-ui/core/Select';
import { EntMedicalType } from '../../api/models/EntMedicalType';
import { EntMedicalEquipment } from '../../api/models/EntMedicalEquipment';

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


const initialSystemequipmentState = {
  //medicalequipment: 'noom',
  noom : '',
 };

interface physiciandata {
  namelist: number;
  
  // create_by: number;
}

function Copyright() {
  return (
    <Typography variant="body2" color="inherit" align="center">
      {'Copyright © '}
      <Link color="inherit" href="https://material-ui.com/">
        Poommin Website
      </Link>{' '}
      {new Date().getFullYear()}
      {'.'}
    </Typography>
  );
}
interface inter {
  nameEquipmentID: number;
  typeEquipmentID: number;
  stockEquipmentID: number;
}
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

    const [systemequipment_data, setSystemequipment_data] = React.useState<
      Partial<inter>
    >({});
  
    const [physicians, setPhysicians] = React.useState<EntPhysician[]>([]);
    const [medicaltypes, setMedicaltypes] = React.useState<EntMedicalType[]>([]);
    const [medicalEquipments, setMedicalEquipments] = React.useState<EntMedicalEquipment[]>([]);
    
  
    const getPhysicians = async () => {
      const res = await api.listPhysician({ limit: 10, offset: 0 });
      setPhysicians(res);
    };
  
    const getMedicaltypes = async () => {
      const res = await api.listMedicaltype({ limit: 10, offset: 0 });
      setMedicaltypes(res);
    };
  
    const getMedicalEquipments = async () => {
      const res = await api.listMedicalequipment({ limit: 10, offset: 0 });
      setMedicalEquipments(res);
    };
  
  
  
    // Lifecycle Hooks
    useEffect(() => {
      getPhysicians();
      getMedicaltypes();
      getMedicalEquipments();
      
    }, []);
  
    // set data to object playlist_video
    const handleChange = (
      event: React.ChangeEvent<{ name?: string; value: unknown }>,
    ) => {
      const name = event.target.name ;
      const { value } = event.target;
      setSystemequipment_data({ ...systemequipment_data, [name]: value });
      console.log(systemequipment_data);
    };
    const handleChange1 = (
      event: React.ChangeEvent<{ name?: string; value: unknown }>,
    ) => {
      const name = event.target.name ;
      const { value } = event.target;
      setSystemequipment_data({ ...systemequipment_data, [name]: value });
      console.log(systemequipment_data);
    };
    const handleChange2 = (
      event: React.ChangeEvent<{ name?: string; value: unknown }>,
    ) => {
      const name = event.target.name ;
      const { value } = event.target;
      setSystemequipment_data({ ...systemequipment_data, [name]: value });
      console.log(systemequipment_data);
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
            
            </Grid>
            
            <Grid item>
                <IconButton 
                style={{ marginLeft: 20 }}
                component={RouterLink}
                to="/login"
                >    
                <HomeIcon color="inherit" />
                </IconButton>
                </Grid>
                <Grid item>
            <Button className={classes.button} variant="outlined" color="inherit" 
            size="small" component={RouterLink}
            to="/">
                logout
              </Button>
                </Grid>  
            <Grid item>
            
            </Grid>
            <Grid item>
              <IconButton color="inherit" className={classes.iconButtonAvatar}>
                <Avatar src="https://shorturl.at/zABQU" alt="NoomAvatar" />
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
            
                </Grid>  
            
            <Grid item>
            
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
              
              <Select
                  name="prefix"
                  value={systemequipment_data.medicalequipmentdata || ''}
                  style={{ width: 200 }}
                  onChange={handleChange}
                >
                  {medicalEquipments.map(item => {
                    return (
                      <MenuItem key={item.id} value={item.id}>
                        {item.name}
                      </MenuItem>
                    )
                  })}
                </Select>
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
            <Select
                  name="types"
                  value={systemequipment_data.typedata || ''}
                  style={{ width: 200 }}
                  onChange={handleChange1}
                >
                  {medicaltypes.map(item => {
                    return (
                      <MenuItem key={item.id} value={item.id}>
                        {item.name}
                      </MenuItem>
                    )
                  })}
                </Select>
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
              
              <Select
                  name="prefix"
                  value={systemequipment_data.medicalequipmentdata || ''}
                  style={{ width: 200 }}
                  onChange={handleChange2}
                >
                  {medicalEquipments.map(item => {
                    return (
                      <MenuItem key={item.id} value={item.id}>
                        {item.stock}
                      </MenuItem>
                    )
                  })}
                </Select>
                  
            </Grid>
            <Grid item xs={2}></Grid>
            <Grid item xs={2}> </Grid>

            <Grid item xs={2}></Grid>
            <Grid item xs={2}></Grid>
            <Grid item xs={2}>
              <Typography color="primary" variant="h6" component="h1">
                วันที่/เวลาที่บันทึก
              </Typography>
            </Grid>
            <Grid item xs={2}>
              
              
                  
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
            <Copyright />
              
            </Grid>
          </Grid>
        </Toolbar>
      </AppBar>
      <ComponanceTable></ComponanceTable>
    </div>
  );
}