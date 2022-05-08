import React, { useState } from 'react';
import { useNavigate, useLocation, Link } from 'react-router-dom';
import AppBar from '@mui/material/AppBar';
import Box from '@mui/material/Box';
import Toolbar from '@mui/material/Toolbar';
import Typography from '@mui/material/Typography';
import Button from '@mui/material/Button';
import IconButton from '@mui/material/IconButton';
import AccountCircle from '@mui/icons-material/AccountCircle';
import MenuItem from '@mui/material/MenuItem';
import Menu from '@mui/material/Menu';
import { useRecoilState, SetterOrUpdater } from 'recoil';
import { userState, defaultUser } from '../config/Recoil';
import { User } from '../config/Type';
import CustomLink from '../atoms/CustomLink';
import LinkStyle from './Header.css';

function Header() {
  const navigate = useNavigate();
  const location = useLocation();
  const [user, setUser]:[User, SetterOrUpdater<User>] = useRecoilState(userState);
  console.log(user);

  const [anchorEl, setAnchorEl] = useState<null | HTMLElement>(null);

  const handleMenu = (event: React.MouseEvent<HTMLElement>) => {
    setAnchorEl(event.currentTarget);
  };

  const handleClose = () => {
    setAnchorEl(null);
  };

  const logout = () => {
    window.localStorage.removeItem('programming-course-market');
    setUser(defaultUser);
    handleClose();
    navigate('/');
  };

  return (
    <Box sx={{ flexGrow: 1 }}>
      <AppBar position="static">
        <Toolbar>
          <Typography variant="h6" component="div" sx={{ flexGrow: 1 }}>
            <CustomLink to="/">
              Skhole
            </CustomLink>
          </Typography>
          {JSON.stringify(user) === JSON.stringify(defaultUser) && (
            <div>
              <CustomLink to="/signup" state={{ from: location }}>
                <Button color="inherit">Signup</Button>
              </CustomLink>
              <CustomLink to="/login" state={{ from: location }}>
                <Button color="inherit">Login</Button>
              </CustomLink>
            </div>
          )}
          {JSON.stringify(user) !== JSON.stringify(defaultUser) && (
            <div>
              <CustomLink to="/course/editor">
                <Button color="inherit">Create Course</Button>
              </CustomLink>
              <IconButton
                size="large"
                aria-label="account of current user"
                aria-controls="menu-appbar"
                aria-haspopup="true"
                onClick={handleMenu}
                color="inherit"
              >
                <AccountCircle />
              </IconButton>
              <Menu
                id="menu-appbar"
                anchorEl={anchorEl}
                anchorOrigin={{
                  vertical: 'top',
                  horizontal: 'right',
                }}
                keepMounted
                transformOrigin={{
                  vertical: 'top',
                  horizontal: 'right',
                }}
                open={Boolean(anchorEl)}
                onClose={handleClose}
              >
                <MenuItem>
                  name:
                  {user.name}
                </MenuItem>
                <MenuItem>
                  email:
                  {user.email}
                </MenuItem>
                <MenuItem>
                  <Link className={LinkStyle} to="/mycontents">
                    MyContents
                  </Link>
                </MenuItem>
                <MenuItem onClick={logout}>Logout</MenuItem>
              </Menu>
            </div>
          )}
        </Toolbar>
      </AppBar>
    </Box>
  );
}

export default Header;
