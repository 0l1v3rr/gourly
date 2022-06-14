const Form = () => {
    return (
        <div className="w-full flex">
            <input type="text" placeholder="Create a new URL" className="border-solid border-teal-400 
                border-2 rounded outline-none text-gray-700 px-2 py-1 w-full" />
            <button type="button" className="ml-2 bg-teal-400 py-1 px-4 text-white rounded border 
                border-teal-300 border-solid leading-none hover:bg-teal-300 transition duration-300">Create</button>
        </div>
    );
}

export default Form;