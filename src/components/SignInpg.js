
import * as React from 'react';
import { useState } from 'react';
import { useNavigate } from 'react-router-dom'; 
import Avatar from '@mui/material/Avatar';
import Button from '@mui/material/Button';
import CssBaseline from '@mui/material/CssBaseline';
import TextField from '@mui/material/TextField';
import FormControlLabel from '@mui/material/FormControlLabel';
import Checkbox from '@mui/material/Checkbox';
import Link from '@mui/material/Link';
import Paper from '@mui/material/Paper';
import Box from '@mui/material/Box';
import Grid from '@mui/material/Grid';
import LockOutlinedIcon from '@mui/icons-material/LockOutlined';
import Typography from '@mui/material/Typography';
import { createTheme, ThemeProvider } from '@mui/material/styles';
import axios from "axios";
import { Form, Modal } from 'react-bootstrap'

function SignInSide() {
  const navigate = useNavigate(); // Initialize useNavigate hook
  const [error, setError] = useState(""); // State to manage error message
  const [addNewUser, setAddNewUser] = useState(false);
  const [newUser, setNewUser] = useState({"handle" : "", "userName": "", "email": "", "subscribedBlogs": [""]});

  const handleSubmit = (event) => {
    event.preventDefault();
    const data = new FormData(event.currentTarget);
    const email = data.get('email');
    const password = data.get('password');

    if ((email === "check@example.com" && password === "password") || (email === "check1@example.com" && password === "PASSWORD")) {
      navigate('/aftersignin'); 
    } else {
      setError("Email or password is incorrect"); 
    }
  };

  const handleSignUpClick = () => {
    navigate('/Signup'); // Navigate to signup page
  };

  const addSingleUser = () => {
    setAddNewUser(false);
    var url = "http://localhost:9000/user/create";
    axios.post(url, {
      "handle" : newUser.handle,
      "userName" : newUser.userName,
      "email" : newUser.email,
      //"subscribedBlogs" : newUser.subscribedBlogs
    }).then(response => {
      if(response.status === 200){
          // Logic for data refresh
      }
    });
  };

  return (
    <ThemeProvider theme={createTheme()}>
      <Grid container component="main" sx={{ height: '100vh' }}>
        <CssBaseline />
        <Grid
          item
          xs={false}
          sm={4}
          md={7}
          sx={{
            backgroundImage: 'url(https://source.unsplash.com/random?wallpapers)',
            backgroundRepeat: 'no-repeat',
            backgroundColor: (t) =>
              t.palette.mode === 'light' ? t.palette.grey[50] : t.palette.grey[900],
            backgroundSize: 'cover',
            backgroundPosition: 'center',
          }}
        />
        <Grid item xs={12} sm={8} md={5} component={Paper} elevation={6} square>
          <Box
            sx={{
              my: 8,
              mx: 4,
              display: 'flex',
              flexDirection: 'column',
              alignItems: 'center',
            }}
          >
            <Avatar sx={{ m: 1, bgcolor: 'secondary.main' }}>
              <LockOutlinedIcon />
            </Avatar>
            <Typography component="h1" variant="h5">
              Sign in
            </Typography>
            <Box component="form" noValidate onSubmit={handleSubmit} sx={{ mt: 1 }}>
              <TextField
                margin="normal"
                required
                fullWidth
                id="email"
                label="Email Address"
                name="email"
                autoComplete="email"
                autoFocus
              />
              <TextField
                margin="normal"
                required
                fullWidth
                name="password"
                label="Password"
                type="password"
                id="password"
                autoComplete="current-password"
              />
              <FormControlLabel
                control={<Checkbox value="remember" color="primary" />}
                label="Remember me"
              />
              <Button
                type="submit"
                fullWidth
                variant="contained"
                sx={{ mt: 3, mb: 2 }}
              >
                Sign In
              </Button>
              {error && <Typography color="error">{error}</Typography>}
              <Grid container>
                <Grid item xs>
                  <Link to="/forgotpassword" variant="body2">
                    Forgot password?
                  </Link>
                </Grid>
                <Grid item>
                  {/* <Button variant="body2" onClick={() => setAddNewUser(true)}>
                    {"Don't have an account? Sign Up"}
                  </Button> */}
                  <Button variant="body2" onClick={() => handleSignUpClick()}>
                    {"Don't have an account? Sign Up"}
                  </Button> 
                  <Modal show={addNewUser} onHide={() => setAddNewUser(false)} centered>
                    <Modal.Header closeButton>
                        <Modal.Title>Sign up</Modal.Title>
                    </Modal.Header>
                    <Modal.Body>
                        <Form.Group>
                            <Form.Label >Codeforces Handle</Form.Label>
                            <Form.Control onChange={(event) => {setNewUser({...newUser, handle: event.target.value})}}/>
                            <Form.Label>Username</Form.Label>
                            <Form.Control onChange={(event) => {setNewUser({...newUser, userName: event.target.value})}}/>
                            <Form.Label >Email Address</Form.Label>
                            <Form.Control onChange={(event) => {setNewUser({...newUser, email: event.target.value})}}/>
                           
                        </Form.Group>
                        <Button onClick={() => addSingleUser()}>Add</Button>
                        <Button onClick={() => setAddNewUser(false)}>Cancel</Button>
                    </Modal.Body>
                  </Modal>
                </Grid>
              </Grid>
            </Box>
          </Box>
        </Grid>
      </Grid>
    </ThemeProvider>
  );
}

export default SignInSide;
