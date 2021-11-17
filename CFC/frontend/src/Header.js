import React, {useState} from 'react'
import { makeStyles } from "@material-ui/core/styles";
import AppBar from "@material-ui/core/AppBar";
import Toolbar from "@material-ui/core/Toolbar";
import Typography from "@material-ui/core/Typography";
import IconButton from "@material-ui/core/IconButton";
import MenuIcon from "@material-ui/icons/Menu";
import green from "@material-ui/core/colors/green";
import { useHistory } from 'react-router';
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
    return (
        <div>
            <AppBar
                color={isCustomColor || isCustomHeight ? "primary" : example}
                className={`${isCustomColor && classes.customColor} ${
                isCustomHeight && classes.customHeight
                }`}
            >
                <Toolbar>
                <IconButton
                    edge="start"
                    className={classes.menuButton}
                    color="inherit"
                    aria-label="menu"
                >
                    <MenuIcon />
                </IconButton>
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
