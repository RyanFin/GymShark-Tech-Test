import React, { useState, useEffect } from 'react';
import { TextField, Button, Box } from '@mui/material';

const ProductForm = ({ onChange, onViewPackSizes }) => {
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
            <Button type="button" variant="contained" color="primary" onClick={() => onViewPackSizes()}>
                View Pack Sizes
            </Button>
        </Box>
    );
};

export default ProductForm;
