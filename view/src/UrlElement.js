const UrlElement = () => {
    return (
        <div className="flex border border-solid border-slate-300 rounded-sm mb-1">
            <a target="_blank" href="http://localhost:8080/url/asdasd" 
                className="border-r border-solid border-slate-300 py-1.5 px-3 
                transition duration-30 hover:text-gray-700">
                localhost:8080/url/asdad
            </a>

            <a target="_blank" href="https://www.google.com"
                className="border-r border-solid border-slate-300 py-1.5 px-3 
                transition duration-30 hover:text-gray-700">
                google.com
            </a>

            <span className="py-1.5 px-3 border-r border-solid border-slate-300">0</span>

            <div className="py-1 px-3">
                <button type="button" className="bg-red-500 py-1.5 px-4 text-white rounded border 
                    border-red-400 border-solid leading-none
                    hover:bg-red-400 transition duration-300"
                >
                    Delete
                </button>
            </div>
        </div>
    );
}

export default UrlElement;