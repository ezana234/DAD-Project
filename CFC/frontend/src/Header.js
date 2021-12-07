import React, {useState} from 'react'
import { makeStyles } from "@material-ui/core/styles";
import AppBar from "@material-ui/core/AppBar";
import Toolbar from "@material-ui/core/Toolbar";
import Typography from "@material-ui/core/Typography";
import IconButton from "@material-ui/core/IconButton";
import MenuIcon from "@material-ui/icons/Menu";
import green from "@material-ui/core/colors/green";
import { useHistory } from 'react-router';
import Menu from '@mui/material/Menu';
import MenuItem from '@mui/material/MenuItem';
//import './Header.css';

const useStyles = makeStyles((theme) => ({
    menuButton: {
      marginRight: theme.spacing(2)
    },
    title: {
      flexGrow: 1
    },
    customColor: {
      // or hex code, this is normal CSS background-color
      backgroundColor: green[500]
    },
    customHeight: {
      minHeight: 200
    },
    offset: theme.mixins.toolbar
  }));

function Header(props) {
  const history = useHistory();
  const classes = useStyles();
  const [example, setExample] = useState("primary");
  const isCustomColor = example === "blue";
  const isCustomHeight = example === "50px";
  const [anchorEl, setAnchorEl] = React.useState(null);
  const open = Boolean(anchorEl);
  const handleClick = (event) => {
    setAnchorEl(event.currentTarget);
  };
  const handleClose = () => {
    setAnchorEl(null);
  };


  const headerClick = () =>{
    if(props.role==1){
      history.push({
        pathname: '/clientHome',
        state: {"Data":props.oldData.Data, "Token":props.oldData.Token, "Role":props.oldData.Role}
    })
    }
    else{
      history.push({
        pathname: '/clinicianHome',
        state: {"Data":props.oldData.Data, "Token":props.oldData.Token, "Role":props.oldData.Role}

    })
    }
  }

    return (
        <div>
            <AppBar
                color={isCustomColor || isCustomHeight ? "primary" : example}
                className={`${isCustomColor && classes.customColor} ${
                isCustomHeight && classes.customHeight
                }`}
            >
                <Toolbar>
                {/* <IconButton
                    edge="start"
                    className={classes.menuButton}
                    color="inherit"
                    aria-label="menu"
                >
                    <MenuIcon
                    id="basic-button"
                    aria-controls="basic-menu"
                    aria-haspopup="true"
                    aria-expanded={open ? 'true' : undefined}
                    onClick={handleClick}
                    >
                      <a>Menu</a>
                    </MenuIcon>
                    <Menu
                    id="basic-menu"
                    anchorEl={anchorEl}
                    open={open}
                    onClose={handleClose}
                    MenuListProps={{
                      'aria-labelledby': 'basic-button',
                    }}
                    onClick={(e) => console.log(e.currentTarget.value)}
                  >
                    <MenuItem value="Home">Home</MenuItem>
                    <MenuItem onClick={(e) => console.log(e.currentTarget.value)}>Profile</MenuItem>
                    {props.role==2 && <MenuItem>Users</MenuItem>}
                    {props.role==2 && <MenuItem>Safety Plan</MenuItem>}
                    <MenuItem onClick={(e) => console.log(e.currentTarget.value)}>Appointments</MenuItem>
                    <MenuItem>Logout</MenuItem>
                  </Menu>
                </IconButton> */}
                <Typography style={{'cursor': 'pointer'}} variant="h6" className={classes.title} onClick={headerClick}>
                    Home
                </Typography>
                <Typography variant="h6" className={classes.title}>
                    {props.header}
                </Typography>
                
                <IconButton color="inherit" onClick={() => history.push("/")}>
                    <h6 style={{paddingTop:"10px"}}>Sign Out</h6>
                </IconButton>
                
                </Toolbar>
            </AppBar>
        </div>
    )
}

export default Header
