import { GetUserLocation } from "./components/getUserLocation";
import React, {useEffect, useState} from "react";
import 'es6-promise/auto'

function App(){  

    const [shop, setShop] = useState('');

    const fetchShop = async () => {
        try{
            const location = await GetUserLocation();
            console.log(location)
            const response = await fetch('http://localhost:8080/api/choiceShop',{
                method: 'POST',
                headers:{
                    'Content-Type':'application/json',
                },
                body: JSON.stringify(location)
            });
            const data = await response.json();
            setShop(data.shop);
        }catch(error){
            console.error('error fetch shop:', error)
        }
    };

    useEffect(()=>{
        const handleSearchClick =()=>{
            fetchShop();
        };

        const searchButton = document.getElementById('searchButton');
        if (searchButton){
            searchButton.addEventListener('click',handleSearchClick);
        }

        return ()=>{
            if(searchButton){
                searchButton.removeEventListener('click',handleSearchClick);
            }
        };
    },[]);

    return(
        <div className="App">
            <header className="App-header">
                <p>今日のお昼はここにしよう</p>
            </header>
            <button id ="searchButton">さがす</button>
            <div id ="printshop">
                {shop && <p>{shop}</p>}
            </div>
        </div>
    );
}

export default App;