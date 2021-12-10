import React from "react";
import './TodoButton.css';

function TodoButton(props) {
    const { children, className, onClick } = props;
  
    return (
      <button className={`todo-button ${className || ""}`} onClick={onClick}>
        <div className="view-my-profile roboto-medium-black-16px">
          {children}
        </div>
      </button>
    );
  }
  
export default TodoButton;