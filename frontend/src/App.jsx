import React, { useState } from 'react';
import axios from 'axios';
import { Container, CssBaseline, Typography } from '@mui/material';
import ProductForm from './components/ProductForm';
import ProductDisplay from './components/ProductDisplay';

const App = () => {
    const [product, setProduct] = useState(null);

    const fetchProductData = async (number) => {
        try {
            const response = await axios.get(`http://localhost:8080/calculate-packs/${number}`);
            setProduct(response.data);
        } catch (error) {
            console.error('Error fetching product data', error);
        }
    };

    return (
        <Container component="main" maxWidth="sm">
            <CssBaseline />
            <Typography variant="h2" component="h1" gutterBottom align="center">
                Gymshark Tech
            </Typography>
            <ProductForm onSubmit={fetchProductData} />
            {product && <ProductDisplay product={product} />}
        </Container>
    );
};

export default App;
