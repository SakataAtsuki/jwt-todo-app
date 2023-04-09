import { Grid, Paper, Typography, TextField, Link, Button, Box } from '@mui/material';
import { signInWithEmailAndPassword } from 'firebase/auth';
import React, { useState } from 'react';
import { useNavigate } from 'react-router-dom';
import { auth } from './firebase';

const Signin = () => {
    const [email, setEmail] = useState('');
    const [password, setPassword] = useState('');
    const navigate = useNavigate();

    const handleChangeEmail = (event) => {
        setEmail(event.target.value);
    }

    const handleChangePassword = (event) => {
        setPassword(event.target.value);
    }

    const handleSubmit = async () => {
        await signInWithEmailAndPassword(auth, email, password)
            .then((userCredential) => {
                const user = userCredential.user;
                user.getIdToken().then(idToken => {
                    localStorage.setItem('jwt', idToken.toString());
                });
                console.log(localStorage.getItem('jwt'));
                navigate("/");
            })
            .catch((err) => {
                alert(err.message);
                console.error(err);
            });
    };

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
                <Typography variant={"h5"} sx={{ m: "30px" }}>Sign in</Typography>
                <TextField
                    label="Email"
                    variant="standard"
                    fullWidth
                    required
                    value={email}
                    onChange={handleChangeEmail} />
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
                        ログイン
                    </Button>
                    <Typography variant="caption">
                        アカウントを持っていませんか？
                        <Link href="/signup">アカウント作成</Link>
                    </Typography>
                </Box>
            </Paper>
        </Grid>
    );
}

export default Signin;
