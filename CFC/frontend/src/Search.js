import React from "react";
import './Search.css'

const SearchBar = ({ searchQuery, setSearchQuery }) => (
    <form action="/" method="get">
        <label htmlFor="header-search">
            <span className="visually-hidden">Search clients</span>
        </label>
        <input
            value={searchQuery}
            onInput={e => setSearchQuery(e.target.value)}
            type="text"
            id="header-search"
            placeholder="Search Clients"
            name="s" 
        />
        <button className="button" type="submit">Search</button>
    </form>
);

export default SearchBar;