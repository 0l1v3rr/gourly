import Form from "./Form";
import UrlElement from "./UrlElement";
import { useState, useEffect } from "react";
import axios from "axios";

const App = () => {
    const [urls, setUrls] = useState([]);

    useEffect(() => {
        axios.get("http://localhost:8080/api/v1/urls")
            .then(res => setUrls(res.data))
            .catch(err => console.log(err));
    }, []);

    const handleFormSubmit = (newUrl) => {
        axios.post("http://localhost:8080/api/v1/urls", { redirect_to: newUrl })
            .then(res => setUrls(current => [...current, res.data]))
            .catch(err => console.log(err));
    };

    const handleUrlDelete = (id) => {
        axios.delete(`http://localhost:8080/api/v1/urls/${id}`)
            .then(_ => {})
            .catch(err => console.log(err));

        const left = [];
        for(const u of [...urls]) {
            if(u.id == id) continue;
            left.push(u);
        }

        setUrls(left);
    };

    return (
        <div className="flex justify-center items-center min-h-screen bg-teal-400">
            <div className="bg-slate-100 shadow-lg p-3 rounded">
                <Form submitBtnEvent={handleFormSubmit} />
                <div className="h-px w-full bg-slate-300 my-3"></div>
                {urls.map(url => <UrlElement key={url.id} url={url} deleteEvent={handleUrlDelete} />)}
            </div>
        </div>
    );
}

export default App;
