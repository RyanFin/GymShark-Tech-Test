import { createTheme } from '@mui/material/styles';

const theme = createTheme({
    palette: {
        primary: {
            main: '#ffffff', // White text
        },
        secondary: {
            main: '#ffffff', // White text
        },
        background: {
            default: '#333333', // Dark grey background
            paper: '#444444', // Slightly lighter grey for paper elements
        },
        text: {
            primary: '#ffffff', // White text
            secondary: '#cccccc', // Light grey text for secondary elements
        },
    },
    typography: {
        fontFamily: 'Roboto, sans-serif',
    },
});

export default theme;
