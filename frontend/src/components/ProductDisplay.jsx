import React from 'react';
import { Box, Typography, Card, CardContent } from '@mui/material';
import { motion } from 'framer-motion';

const ProductDisplay = ({ product }) => {
    return (
        <motion.div
            initial={{ opacity: 0, y: -50 }}
            animate={{ opacity: 1, y: 0 }}
            transition={{ duration: 0.5 }}
        >
            <Card sx={{ maxWidth: 600, margin: 'auto', mt: 4 }}>
                <CardContent>
                    <Typography variant="h5" component="div">
                        {product.name}
                    </Typography>
                    <Typography variant="body2" color="text.secondary">
                        Price: £{product.price}
                    </Typography>
                    <Box mt={2}>
                        {Object.entries(product.packs).map(([size, quantity]) => (
                            <Typography key={size} variant="body2" color="text.secondary">
                                Pack Size: {size}, Quantity: {quantity}
                            </Typography>
                        ))}
                    </Box>
                </CardContent>
            </Card>
        </motion.div>
    );
};

export default ProductDisplay;
