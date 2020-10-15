import React, { useState, useEffect } from 'react';
import { makeStyles } from '@material-ui/core/styles';
import Table from '@material-ui/core/Table';
import TableBody from '@material-ui/core/TableBody';
import TableCell from '@material-ui/core/TableCell';
import TableContainer from '@material-ui/core/TableContainer';
import TableHead from '@material-ui/core/TableHead';
import TableRow from '@material-ui/core/TableRow';
import Paper from '@material-ui/core/Paper';
import Button from '@material-ui/core/Button';
import { DefaultApi } from '../../api/apis';
import { EntSystemequipment } from '../../api/models/EntSystemequipment';
 
import moment from 'moment';
const useStyles = makeStyles({
 table: {
   minWidth: 650,
 },
});
 
export default function ComponentsTable() {
 const classes = useStyles();
 const api = new DefaultApi();

 const [systemequipments, setSystemequipments] = useState<EntSystemequipment[]>();
 const [loading, setLoading] = useState(true);
 
 useEffect(() => {
   const getSystemequipments = async () => {
     const res = await api.listSystemequipment({ limit: 10, offset: 0 });
     setLoading(false);
     setSystemequipments(res);
   };
   getSystemequipments();
 }, [loading]);
 
 const deleteSystemequipments = async (id: number) => {
   const res = await api.deleteSystemequipment({ id: id });
   setLoading(true);
 };
 console.log(systemequipments)
 return (
   <TableContainer component={Paper}>
     <Table className={classes.table} aria-label="simple table">
       <TableHead>
         <TableRow>
         <TableCell align="center">No.</TableCell>
         <TableCell align="center">Medical</TableCell>
         <TableCell align="center">Type</TableCell>
         <TableCell align="center">Stock</TableCell>
         <TableCell align="center">PHYSICIAN</TableCell>
         <TableCell align="center">DATA</TableCell>
         <TableCell align="center">Manage</TableCell>
         </TableRow>
       </TableHead>
       <TableBody>
         {systemequipments === undefined 
          ? null
          : systemequipments.map((item :any)=> (
           <TableRow key={item.id}>
             <TableCell align="center">{item.id}</TableCell>
             <TableCell align="center">{item.edges?.medicalequipment?.name}</TableCell>
             <TableCell align="center">{item.edges?.medicaltype?.name}</TableCell>
             <TableCell align="center">{item.stock}</TableCell>
             <TableCell align="center">{item.edges?.physician?.name}</TableCell>
             <TableCell align="center">{moment(item.addedtime).format('DD/MM/YYYY HH:mm:ss')}</TableCell>
            
             <TableCell align="center">
               <Button
                 onClick={() => {
                   deleteSystemequipments(item.id);
                 }}
                 style={{ marginLeft: 10 }}
                 variant="contained"
                 color="secondary"
               >
                 Delete
               </Button>
             </TableCell>
           </TableRow>
         ))}
       </TableBody>
     </Table>
   </TableContainer>
 );
}
