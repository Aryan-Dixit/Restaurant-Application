import React, { Component } from 'react';
import axios from "axios";

class RestaurantUpationbyID extends Component {
  initialState = {
    id :"",
      name: "",
      address: "",
      addressline2: "",
      url: "",
      outcode: "",
      postcode: "",
      rating: "",
      type_of_food: "",

      errors: [],
    restaurantList:[]
  };

  state = this.initialState;
  FindRestaurantByID=()=>{
    //const {id } = this.state;
    axios.get("http://localhost:8080/restaurant/"+this.state.id).then(response => {
      this.setState({restaurantList : response.data.responseData,
        name : response.data.responseData.name,
        address : response.data.responseData.address,
        addressline2 :response.data.responseData.addressline2,
        url : response.data.responseData.url,
        outcode : response.data.responseData.outcode,
        postcode : response.data.responseData.postcode,
        rating : response.data.responseData.rating,
        type_of_food : response.data.responseData.type_Of_food
      })
    })
  }
  
    UpdateRestaurantByID=()=>{
      axios.put('http://localhost:8080/restaurant/'+this.state.id, 
      {
        name: this.state.name, 
        address: this.state.address, 
        addressline2: this.state.addressline2,
        url:this.state.url,
        outcode:this.state.outcode,
        postcode:this.state.postcode,
        rating:parseInt(this.state.rating, 10),
        type_of_food:this.state.type_of_food

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
    type_of_food: ''
  });
  alert("Restaurant Data Updated")
    }
 
  
  render() {
    return (
      <div className="FormCenter">
        <form className="FormFields">
        <div className="FormField">
        
          <div class="ui form">
    <div class="required eight wide field">
  <h style={{ color: 'White' }}>Enter Restaurant Id</h>

        <div class="ui action input">
        <input
          value={this.state.id}
          onChange={evt => this.setState({ id: evt.target.value })}
          type="text"
          className="FormField__Input"
          placeholder="Enter Restaurant ID"
        />
      
          <button 
                type="button" 
                className="ui button"
                onClick={()=>{
                console.log("click post");
                this.FindRestaurantByID()}}>Find Restaurant</button>
        </div>      
     </div>  
</div>
</div>
        
       <div className="FormField">
       <h1>Edit Fields to Update Restaurant
       </h1>
       <br></br>
                <label className="FormField__Label" htmlFor="name">Restaurant Name</label>
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
          value={this.state.type_of_food}
          onChange={evt => this.setState({ type_of_food: evt.target.value })}
          type="text"
          className="FormField__Input"
          placeholder="Type Of Food"
        />
        </div>  
        <button type="button"
        className="FormField__Button mr-20" 
                onClick={()=>{
                console.log("click post");
                this.UpdateRestaurantByID()}}>Update Restaurant</button>
      </form>
      </div>
    );
  }
}

export default RestaurantUpationbyID;
