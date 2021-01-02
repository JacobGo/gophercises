import './style';
import { Component } from 'preact';

export default class App extends Component {
	state = {
		mode: '0',
		n: '1',
		image: null,
		loading: false,
		processedImage: ''
	}

	onSubmit = async e => {
		e.preventDefault()
		this.setState({...this.state, loading: true})
		const data = new FormData()
		Object.keys(this.state).forEach(key => data.append(key, this.state[key]))
		const response = await fetch('/upload', {
			method: 'POST',
			body: data
		})

		const image = await response.text()
		this.setState({...this.state, processedImage: image, loading: false})
	}

	onModeInput = e => {
		this.setState({...this.state, mode: e.target.value})
		e.preventDefault()
	}

	onNInput = async e => {
		await this.setState({...this.state, n: e.target.value})
		await this.onSubmit(e)
		e.preventDefault()
	}

	onImageInput = e => {
		this.setState({...this.state, image: e.target.files[0]})
		e.preventDefault()
	}

	modes = ["Combo", "Triangle", "Rectangle", "Ellipse", "Circle", "Rotated Rectangle", "Beziers", "Rotated Ellipse", "Polygon"]

	render() {
		return (
			<div>
				<h1>Image Transformer</h1>
				<p>Upload an image to be reproduced with simple geometric shapes through <a href="https://primitive.lol/">primitive</a>!</p>
				<form onsubmit={this.onSubmit}>
					<div>
						<label for="image">Image: </label>
						<input type="file" onChange={this.onImageInput} required/>
					</div>
					<div>
						<label for="mode">Mode: </label>
						<select id="mode" onChange={this.onModeInput} value={this.state.mode} required>
							{this.modes.map((mode, i) => {
								return <option value={i}>{mode}</option>
							})}
						</select>
					</div>

					<div>
						<label for="n">Number of Shapes: </label>
						<input type="number" id="n" min="1" max="200" onChange={this.onNInput} value={this.state.n} required/>
					</div>

					<div class="submitButton">
						<button type="submit">Transform!</button>
					</div>

					{this.state.loading
						? <p>Loading...</p>
						: <></>
					}

				</form>

				<div class="generatedImage" dangerouslySetInnerHTML={{__html: this.state.processedImage}} />
			</div>
		);
	}
}
