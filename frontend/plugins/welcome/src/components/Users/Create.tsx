import React, { useState } from 'react';

import { Link as RouterLink } from 'react-router-dom';

import {

 Content,

 Header,

 Page,

 pageTheme,

 ContentHeader,

} from '@backstage/core';

import { makeStyles, Theme, createStyles } from '@material-ui/core/styles';

import TextField from '@material-ui/core/TextField';

import Button from '@material-ui/core/Button';

import FormControl from '@material-ui/core/FormControl';

import { Alert } from '@material-ui/lab';

import { DefaultApi } from '../../api/apis';

import Avatar from '@material-ui/core/Avatar';

import Autocomplete from '@material-ui/lab/Autocomplete';

import Typography from '@material-ui/core/Typography';

import Box from '@material-ui/core/Box';

import Link from '@material-ui/core/Link';

const useStyles = makeStyles((theme: Theme) =>

 createStyles({

   root: {

     display: 'flex',

     flexWrap: 'wrap',

     justifyContent: 'center',

   },

   margin: {

     margin: theme.spacing(3),

   },

   withoutLabel: {

     marginTop: theme.spacing(3),

   },

   textField: {

     width: '25ch',

   },

 }),

);


const initialUserState = {

 name: 'System Analysis and Design',

 age: 20,

};


export default function Create() {

 const classes = useStyles();

 const profile = { givenName: 'ระบบเครื่องมือแพทย์' };

 const api = new DefaultApi();


 const [user, setUser] = useState(initialUserState);

 const [status, setStatus] = useState(false);

 const [alert, setAlert] = useState(true);


 const handleInputChange = (event: any) => {

   const { id, value } = event.target;

   setUser({ ...user, [id]: value });

 };


 const CreateUser = async () => {

   const res = await api.createUser({ user });

   setStatus(true);

   if (res.id != ''){

     setAlert(true);

   } else {

     setAlert(false);

   }

   const timer = setTimeout(() => {

     setStatus(false);

   }, 1000);

 };

 const frist = [
  { title: '10001' },
  { title: '10002'},];

 return (

   <Page theme={pageTheme.home}>

<Header

title={` ${profile.givenName || 'ระบบเครื่องมือแพทย์'}`}

subtitle="ยินดีต้อนรับเข้าสู่ระบบบันทึกข้อมูลอุปกรณ์แพทย์"
><Avatar>N</Avatar>
<Typography component="div" variant="body1">
      <Box color="white">Poommin Phinphimai</Box>
      <Box color="secondary.main"></Box>
      
    
    <Typography className={classes.root}>
      <Link href="#" onClick={"link"}>
        Log Out
      </Link>
      </Typography>
      </Typography>
</Header>

     <Content>

       <ContentHeader title="">

         {status ? (

           <div>

             {alert ? (

               <Alert severity="success">

                 This is a success alert — check it out!

               </Alert>

             ) : (

               <Alert severity="warning" style={{ marginTop: 20 }}>

                 This is a warning alert — check it out!

               </Alert>

             )}

           </div>

         ) : null}

       </ContentHeader>

       <div className={classes.root}>

         <form noValidate autoComplete="off">

         <div className={classes.margin}>
         <div class="form-control">
          <label for="my-input">ID เครื่องมือแพทย์     </label>
          <input id="my-input" aria-describedby="my-helper-text" />
          <span id="my-helper-text"></span>
          </div>

        <div class="form-control">
          <label for="my-input">ชื่อเครื่องมือแพทย์     </label>
          <input id="my-input" aria-describedby="my-helper-text" />
          <span id="my-helper-text"></span>
        </div>

        <div class="form-control">
          <label for="my-input">จำนวน/ชิ้น&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;</label>
          <input id="my-input" aria-describedby="my-helper-text" />
          <span id="my-helper-text"></span>
        </div>

         </div>

           <div className={classes.margin}>

             <Button

               onClick={() => {

                 CreateUser();

               }}

               variant="contained"

               color="primary"

             >

              เพิ่มข้อมูล

             </Button>

             <Button

               style={{ marginLeft: 20 }}

               component={RouterLink}

               to="/"

               variant="contained"

             >

              อัพเดต

             </Button>
             <Button

               style={{ marginLeft: 20 }}

               component={RouterLink}

               to="/"

               variant="contained"

               color="secondary"

             >

              ลบข้อมูล

             </Button>
             <h1>ระบบทำรายการ</h1>
             <Autocomplete
            id="combo-box-demo"
            options={frist}
            getOptionLabel={(option) => option.title}
            style={{ width: 400,marginTop: 10 }}
            renderInput={(params) => <TextField {...params} label="ID เครื่องมือแพทย์" variant="outlined" 
            />}
            />

            <Autocomplete
            id="combo-box-demo"
            options={frist}
            getOptionLabel={(option) => option.title}
            style={{ width: 400,marginTop: 10 }}
            renderInput={(params) => <TextField {...params} label="จำนวน/ชิ้น" variant="outlined" 
            />}
            />
            
           

           </div>
           <div className={classes.root}>
           <Button

            onClick={() => {

              CreateUser();

            }}

            variant="contained"

            color="primary"

            >

            บันทึกข้อมูล

            </Button>

              
           </div>
           
        
         </form>

       </div>

     </Content>

   </Page>

 );

}