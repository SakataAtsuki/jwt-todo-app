import { AppBar, Box, Button, Toolbar, Typography } from "@mui/material";
import { useNavigate } from 'react-router-dom';
import { auth } from "../auth/firebase";

const Header = () => {
    const navigate = useNavigate();

    const handleLogout = () => {
        auth.signOut()
            .then(() => {
                localStorage.removeItem('jwt');
                navigate("/signin");
            })
            .catch((err) => {
                alert(err.message);
                console.error(err);
            })
    };

    return (
        <Box sx={{ flexGrow: 1 }}>
            <AppBar position="static">
                <Toolbar>
                    <Typography variant="h6" component="div" sx={{ flexGrow: 1 }}>
                        Task Manager
                    </Typography>
                    <Button color="inherit" onClick={handleLogout}>ログアウト</Button>
                </Toolbar>
            </AppBar>
        </Box>
    );
}

export default Header;
