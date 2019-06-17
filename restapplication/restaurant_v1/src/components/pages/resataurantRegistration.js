import React from "react";
import axios from "axios";


function validate(name, address, addressline2,url,outcode,postcode,rating,typeoffood ) {

  const errors = [];

  if (name.length === 0) {
    errors.push("Name can't be empty");
  }
  if (address.length === 0) {
    errors.push("Address can't be empty");
  }
  if (addressline2.length === 0) {
    errors.push("Address Line 2 can't be empty");
  }
  if (url.length === 0) {
    errors.push("URL can't be empty");
  }
  if (outcode.length === 0) {
    errors.push("Outcode can't be empty");
  }
  if (postcode.length === 0) {
    errors.push("Postcode can't be empty");
  }
  if (rating.length === 0) {
    errors.push("Rating can't be empty");
  }
  if (typeoffood.length === 0) {
    errors.push("Type Of Food can't be empty");
  }

  return errors;
}

class restaurantRegistration extends React.Component {
  constructor() {
    super();
    this.state = {
      name: "",
      address: "",
      addressline2: "",
      url: "",
      outcode: "",
      postcode: "",
      rating: "",
      typeoffood: "",

      errors: []
    };

    this.handleSubmit = this.handleSubmit.bind(this);
  }

  handleSubmit(e) {
    e.preventDefault();

    const { name, address, addressline2,url,outcode,postcode,rating,typeoffood } = this.state;

    const errors = validate(name, address, addressline2,url,outcode,postcode,rating,typeoffood );
    if (errors.length > 0) {
      this.setState({ errors });
      return;
    }
        axios.post('http://localhost:8080/restaurant/', 
        {
          name: this.state.name, 
          address: this.state.address, 
          addressline2: this.state.addressline2,
          url:this.state.url,
          outcode:this.state.outcode,
          postcode:this.state.postcode,
          rating:parseInt(this.state.rating, 10),
          type_of_food:this.state.typeoffood

        }, 
        {header: {
          'Access-Control-Allow-Origin': '*',
          'Content-Type': 'application/json',
      }},
        )
        .then(function (response) {
            console.log(response);
        })
        .catch(function (error) {
            console.log(error);
    });
        this.setState({
          name: '',
          address: '',
          addressline2: '',
          url: '',
          outcode: '',
          postcode: '',
          rating: '',
          typeoffood: ''
        });
        alert("Restaurant Data Submited")
  }

  render() {
    const { errors } = this.state;
    return (
      <div className="FormCenter">
      <form onSubmit={this.handleSubmit} className="FormFields">
        {errors.map(error => (
          <p key={error}>Error: {error}</p>
        ))}
        <div className="FormField">
                <label className="FormField__Label" htmlFor="name">Full Name</label>
        <input
          value={this.state.name}
          onChange={evt => this.setState({ name: evt.target.value })}
          type="text"
          className="FormField__Input"
          placeholder="Name"
        />
        </div>

        <div className="FormField">
                <label className="FormField__Label" htmlFor="address">Address</label>
      
        <input
          value={this.state.address}
          onChange={evt => this.setState({ address: evt.target.value })}
          type="text"
          className="FormField__Input"
          placeholder="Address"
        />
        </div>

<div className="FormField">
                <label className="FormField__Label" htmlFor="addressline2">Address Line 2</label>
      
        <input
          value={this.state.addressline2}
          onChange={evt => this.setState({ addressline2: evt.target.value })}
          type="text"
          className="FormField__Input"
          placeholder="Address Line 2"
        />
        </div>

<div className="FormField">
                <label className="FormField__Label" htmlFor="url">Restaurant URL</label>
      
        <input
          value={this.state.url}
          onChange={evt => this.setState({ url: evt.target.value })}
          type="text"
          className="FormField__Input"
          placeholder="URL"
        />
        </div>
        <div className="FormField">
                <label className="FormField__Label" htmlFor="outcode">Outcode</label>
      
        <input
          value={this.state.outcode}
          onChange={evt => this.setState({ outcode: evt.target.value })}
          type="text"
          className="FormField__Input"
          placeholder="Outcode"
        />
        </div>
        <div className="FormField">
                <label className="FormField__Label" htmlFor="postcode">Postcode</label>
      
        <input
          value={this.state.postcode}
          onChange={evt => this.setState({ postcode: evt.target.value })}
          type="text"
          className="FormField__Input"
          placeholder="Postcode"
        /></div>
        <div className="FormField">
                <label className="FormField__Label" htmlFor="rating">Rating</label>
      
        <input
          value={this.state.rating}
          onChange={evt => this.setState({ rating: evt.target.value })}
          type="number"
          className="FormField__Input"
          placeholder="Rating"
        />
        </div>
        <div className="FormField">
                <label className="FormField__Label" htmlFor="typeoffood">Type Of Food</label>
      
        <input
          value={this.state.typeoffood}
          defaultValue="hello"
          onChange={evt => this.setState({ typeoffood: evt.target.value })}
          type="text"
          className="FormField__Input"
          placeholder="Type Of Food"
        />
        </div>  
        <button type="submit" className="FormField__Button mr-20">Register</button>
      </form>
      </div>
    );
  }
}

export default restaurantRegistration;