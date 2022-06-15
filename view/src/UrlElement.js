import axios from "axios";
import { useState } from "react";

const UrlElement = (props) => {
    const [url, setUrl] = useState(props.url);

    const handleDeleteClick = () => {
        props.deleteEvent(url.id);
    };

    const handleIncrementView = (e) => {
        e.preventDefault();
        window.open(`http://localhost:8080/url/${url.unique_id}`, '_blank', 'noopener,noreferrer');

        let incremented = {...url};
        incremented.clicks++;
        setUrl(incremented);

        axios.patch(`http://localhost:8080/api/v1/urls/${url.id}`)
            .then(_ => {})
            .catch(err => console.log(err));
    };
    
    return (
        <div className="flex border border-solid border-slate-300 rounded-sm mb-2 w-full">
            <a target="_blank" href={`http://localhost:8080/url/${url.unique_id}`} 
                onClick={handleIncrementView} 
                className="border-r border-solid border-slate-300 py-1.5 px-3 
                transition duration-30 hover:text-gray-600 hover:underline w-64 text-center">
                localhost:8080/url/{url.unique_id}
            </a>

            <a target="_blank" href={url.redirect_to}
                className="border-r border-solid border-slate-300 py-1.5 px-3 
                transition duration-30 hover:text-gray-600 hover:underline w-80 text-center">
                {url.redirect_to}
            </a>

            <span className="py-1.5 px-3 border-r border-solid border-slate-300 w-10 text-center">
                {url.clicks}
            </span>

            <div className="py-1 px-3">
                <button type="button" onClick={handleDeleteClick} className="bg-red-500 py-1.5 
                    px-4 text-white rounded border  border-red-400 border-solid leading-none
                    hover:bg-red-400 transition duration-300">
                    Delete
                </button>
            </div>
        </div>
    );
}

export default UrlElement;