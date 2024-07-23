import React, { useState, useEffect } from 'react';
import { TextField, Box } from '@mui/material';

const ProductForm = ({ onChange }) => {
    const [number, setNumber] = useState('');

    useEffect(() => {
        if (number) {
            onChange(number);
        }
    }, [number, onChange]);

    return (
        <Box
            component="div"
            sx={{
                display: 'flex',
                flexDirection: 'column',
                alignItems: 'center',
                gap: 2,
                color: 'text.primary'
            }}
        >
            <TextField
                label="Enter Number"
                variant="outlined"
                value={number}
                onChange={(e) => setNumber(e.target.value)}
                InputLabelProps={{ style: { color: '#ffffff' } }}
                InputProps={{ style: { color: '#ffffff' } }}
            />
        </Box>
    );
};

export default ProductForm;
