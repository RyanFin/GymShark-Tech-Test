import React from 'react';
import { Box, Typography, List, ListItem, Card, CardContent } from '@mui/material';
import { motion } from 'framer-motion';

const PackSizesDisplay = ({ packSizes }) => {
    return (
        <motion.div
            initial={{ opacity: 0, y: -50 }}
            animate={{ opacity: 1, y: 0 }}
            transition={{ duration: 0.5 }}
        >
            <Card sx={{ maxWidth: 600, margin: 'auto', mt: 4, backgroundColor: 'background.paper', color: 'text.primary' }}>
                <CardContent>
                    <Typography variant="h5" component="div">
                        Pack Sizes
                    </Typography>
                    <List>
                        {packSizes.map((size, index) => (
                            <motion.li
                                key={index}
                                initial={{ opacity: 0, x: -50 }}
                                animate={{ opacity: 1, x: 0 }}
                                transition={{ duration: 0.3, delay: index * 0.1 }}
                            >
                                <ListItem>
                                    <Typography variant="body2" color="text.secondary">
                                        {size}
                                    </Typography>
                                </ListItem>
                            </motion.li>
                        ))}
                    </List>
                </CardContent>
            </Card>
        </motion.div>
    );
};

export default PackSizesDisplay;
