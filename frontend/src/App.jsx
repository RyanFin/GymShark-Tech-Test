import React, { useState, useCallback } from 'react';
import axios from 'axios';
import { Container, CssBaseline, Typography, Box } from '@mui/material';
import ProductForm from './components/ProductForm';
import ProductDisplay from './components/ProductDisplay';
import gymsharkLogo from '/gymshark-logo.png';

const App = () => {
    const [product, setProduct] = useState(null);

    const fetchProductData = useCallback(async (number) => {
        try {
            const response = await axios.get(`http://localhost:8080/calculate-packs/${number}`);
            setProduct(response.data);
        } catch (error) {
            console.error('Error fetching product data', error);
        }
    }, []);

    return (
        <Container
            component="main"
            maxWidth="sm"
            sx={{
                backgroundColor: 'background.default',
                minHeight: '100vh',
                padding: 4,
                color: 'text.primary'
            }}
        >
            <CssBaseline />
            <Box display="flex" alignItems="center" justifyContent="center" mt={4} mb={4}>
                <img src={gymsharkLogo} alt="Gymshark Logo" style={{ width: 50, height: 50, marginRight: 16 }} />
                <Typography variant="h2" component="h1" gutterBottom align="center" color="text.primary">
                    Gymshark Tech
                </Typography>
            </Box>
            <ProductForm onChange={fetchProductData} />
            {product && <ProductDisplay product={product} />}
        </Container>
    );
};

export default App;
