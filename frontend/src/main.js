import { ready } from '@wails/runtime'
import App from './App.svelte';

let app;

ready(() => {
	app = new App({
		target: document.body,
	});
});

export default app;