import Form from "./Form";
import UrlElement from "./UrlElement";

const App = () => {
	return (
		<div className="flex justify-center items-center min-h-screen bg-teal-400">
			<div className="bg-slate-100 shadow-lg p-3 rounded">
				<Form />
				<div className="h-px w-full bg-slate-300 my-3"></div>
				<UrlElement />
			</div>
		</div>
	);
}

export default App;
