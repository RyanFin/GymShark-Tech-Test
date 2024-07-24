import React, { useState, useCallback } from 'react';
import axios from 'axios';
import { Container, CssBaseline, Typography, Box, Snackbar, Alert } from '@mui/material';
import ProductForm from './components/ProductForm';
import ProductDisplay from './components/ProductDisplay';
import PackSizesDisplay from './components/PackSizesDisplay';
import gymsharkLogo from '/gymshark-logo.png';

const App = () => {
    const [product, setProduct] = useState(null);
    const [packSizes, setPackSizes] = useState([]);
    const [snackbarMessage, setSnackbarMessage] = useState('');
    const [snackbarSeverity, setSnackbarSeverity] = useState('success');
    const [snackbarOpen, setSnackbarOpen] = useState(false);

    const backendUrl = process.env.REACT_APP_BACKEND_URL;
    // const backendUrl = "http://localhost:8080";

    const fetchProductData = useCallback(async (number) => {
        try {
            const response = await axios.get(`${backendUrl}/calculate-packs/${number}`);
            setProduct(response.data);
        } catch (error) {
            console.error('Error fetching product data', error);
        }
    }, []);

    const fetchPackSizes = useCallback(async () => {
        try {
            const response = await axios.get(`${backendUrl}/view-packsizes`);
            setPackSizes(response.data.packSizes);
        } catch (error) {
            console.error('Error fetching pack sizes', error);
        }
    }, []);

    const handleAddPackSize = useCallback(async (packsize) => {
        try {
            const response = await axios.post(`${backendUrl}add-packsize?packsize=${packsize}`);
            setSnackbarMessage(`Successfully added pack size ${packsize}`);
            setSnackbarSeverity('success');
            setSnackbarOpen(true);
            fetchPackSizes();
        } catch (error) {
            setSnackbarMessage(`Error adding pack size ${packsize}`);
            setSnackbarSeverity('error');
            setSnackbarOpen(true);
        }
    }, [fetchPackSizes]);

    const handleRemovePackSize = useCallback(async (packsize) => {
        try {
            const response = await axios.delete(`${backendUrl}/remove-packsize?packsize=${packsize}`);
            setSnackbarMessage(`Successfully removed pack size ${packsize}`);
            setSnackbarSeverity('success');
            setSnackbarOpen(true);
            fetchPackSizes();
        } catch (error) {
            setSnackbarMessage(`Error removing pack size ${packsize}`);
            setSnackbarSeverity('error');
            setSnackbarOpen(true);
        }
    }, [fetchPackSizes]);

    const handleSnackbarClose = () => {
        setSnackbarOpen(false);
    };

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
            <ProductForm 
                onChange={fetchProductData} 
                onViewPackSizes={fetchPackSizes} 
                onAddPackSize={handleAddPackSize} 
                onRemovePackSize={handleRemovePackSize} 
            />
            {product && <ProductDisplay product={product} />}
            {packSizes.length > 0 && <PackSizesDisplay packSizes={packSizes} />}
            <Snackbar
                open={snackbarOpen}
                autoHideDuration={6000}
                onClose={handleSnackbarClose}
            >
                <Alert onClose={handleSnackbarClose} severity={snackbarSeverity}>
                    {snackbarMessage}
                </Alert>
            </Snackbar>
        </Container>
    );
};

export default App;
