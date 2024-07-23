import React, { useState } from 'react';
import { TextField, Button, Box } from '@mui/material';

const ProductForm = ({ onSubmit }) => {
    const [number, setNumber] = useState('');

    const handleSubmit = (e) => {
        e.preventDefault();
        onSubmit(number);
    };

    return (
        <Box component="form" onSubmit={handleSubmit} sx={{ display: 'flex', flexDirection: 'column', alignItems: 'center', gap: 2 }}>
            <TextField 
                label="Enter Number" 
                variant="outlined" 
                value={number}
                onChange={(e) => setNumber(e.target.value)}
            />
            <Button type="submit" variant="contained">Submit</Button>
        </Box>
    );
};

export default ProductForm;
