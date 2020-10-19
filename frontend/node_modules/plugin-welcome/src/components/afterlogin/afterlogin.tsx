import React, { FC } from 'react';
import { Link as RouterLink } from 'react-router-dom';
import ComponanceTable from '../Table';
import Button from '@material-ui/core/Button';
 
import {
 Content,
 Header,
 Page,
 pageTheme,
 ContentHeader,
 Link,
} from '@backstage/core';
 
const afterlogin: FC<{}> = () => {
 const profile = { givenName: 'to ระบบเครื่องมือแพทย์' };
 
 return (
   <Page theme={pageTheme.home}>
     <Header
       title={`Welcome ${profile.givenName || 'to Backstage'}`}
       subtitle="ระบบหลัก : ระบบข้อมูลแพทย์"
     ></Header>
     <Content>
       <ContentHeader title="บันทึกข้อมูล">
         <Link component={RouterLink} to="/user">
           <Button variant="contained" color="primary">
             Add User
           </Button>
         </Link>
       </ContentHeader>
       <ComponanceTable></ComponanceTable>
     </Content>
     
   </Page>
 );
};
 
export default afterlogin;