import { Grid, Paper, Typography, TextField, Link, Button, Box } from '@mui/material';
import React, { useState } from 'react';
import { useNavigate } from 'react-router-dom';
import { createUserWithEmailAndPassword } from 'firebase/auth';
import { auth } from './firebase';

const Signup = () => {
    const navigate = useNavigate();
    const [email, setEmail] = useState('');
    const [password, setPassword] = useState('');

    const handleChangeEmail = (event) => {
        setEmail(event.target.value);
    }

    const handleChangePassword = (event) => {
        setPassword(event.target.value);
    }

    const handleSubmit = async () => {
        await createUserWithEmailAndPassword(auth, email, password)
            .then(() => {
                navigate("/");
            }).catch((err) => {
                alert(err.message);
                console.error(err);
            });
    }

    return (
        <Grid>
            <Paper
                elevation={3}
                sx={{
                    p: 4,
                    height: "70vh",
                    width: "50vh",
                    m: " 20px auto"
                }}
            >
                <Typography variant={"h5"} sx={{ m: "30px" }}>Sign up</Typography>
                <TextField
                    label="Email"
                    variant="standard"
                    fullWidth
                    required
                    value={email}
                    onChange={handleChangeEmail}
                />
                <TextField
                    type="password"
                    label="Password"
                    variant="standard"
                    fullWidth
                    value={password}
                    onChange={handleChangePassword}
                    required />
                <Box mt={4}>
                    <Button type="submit" color="primary" variant="contained" fullWidth onClick={handleSubmit}>
                        登録
                    </Button>
                    <Typography variant="caption">
                        <Link href="/signin">アカウントをお持ちですか？</Link>
                    </Typography>
                </Box>
            </Paper>
        </Grid>
    );
}

export default Signup;
