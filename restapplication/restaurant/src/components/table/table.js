import React, { Component } from 'react';
import { Wrapper, AgroTable, Th } from "./table.style";
import faker from 'faker';
import axios from "axios";
//import { Input } from "../input/input.style"

class Table extends Component {
  initialState = {
    searchTerm: '',
    RestaurantsList:[],
    restaurantList:[]
  };

  state = this.initialState;
   componentDidMount() {
    axios.get("http://localhost:8080/restaurant/").then(response => {
      this.setState({restaurantList : response.data.responseData.restaurants})
    })

  } 
  onSearch = ev => {
    let RestaurantsList;
    const searchTerm = ev.target.value;
    if (!searchTerm) {
      RestaurantsList = this.state.restaurantList;
    } else {
      RestaurantsList = this.state.RestaurantsList.filter(item => {
        return item.name.toLowerCase().includes(searchTerm.toLowerCase());
      })
    }
    this.setState({ searchTerm: searchTerm, RestaurantsList });
  }

  renderList = lists => {
    return lists.map(list => {
      return (
        <tr>
          <td>{list.id}</td>
          <td>{list.name}</td>
          <td>{list.address}</td>
          <td>{list.addressline2}</td>
          <td>{list.outcode}</td>
          <td>{list.postcode}</td>
          <td>{list.rating}</td>
          <td>{list.type_of_food}</td>
      
          <td>{
          <div class="ui card">
            <div class="image">
              <a href={list.url} class="ui medium image">
                <img src={faker.image.avatar()} alt="food"/>
              </a>
            </div>
            <div class="content">
              <div class="description">
                {list.name}
              </div>
            </div>
          </div>
        }
        </td>
        </tr>
      )
    })
  }

  render() {
    return (
      <>
      <div class="ui search">
  <div class="ui icon input">
  <input type="text" placeholder="search by name" onChange={this.onSearch} />
    <i class="search icon"></i>
  </div>
  </div>
               <Wrapper>
          <AgroTable>
            <thead>
              <tr>
                <Th>Restaurant Id</Th>
                <Th>Name</Th>
                <Th>address</Th>
                <Th>addressLine2</Th>
                <Th>outcode</Th>
                <Th>postcode</Th>
                <Th>rating</Th>
                <Th>type_of_food</Th>
                <Th>Card</Th>
              </tr>
            </thead>
            <tbody>
              {this.renderList(this.state.restaurantList)}
            </tbody>
          </AgroTable>
        </Wrapper>
      </>
    );
  }
}

export default Table;
