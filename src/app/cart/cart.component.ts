import { Component } from '@angular/core';
import { Router } from '@angular/router';
import { HttpService } from '../http.service';
import * as http from '../http.service';

@Component({
  selector: 'app-cart',
  templateUrl: './cart.component.html',
  styleUrls: ['./cart.component.css']
})
export class CartComponent {
  products : http.ItemList;
  displayedColumns : any[] = ['name','imgUrl','price','rating','store_id', 'remove_button'];

  constructor (private router : Router, private httpService: HttpService) {}

  // Routing function to take user to Home Page
  goToPage(pageName:string):void {
    this.router.navigate([`${pageName}`]);
  }

  ngOnInit() {
    this.httpService.getProducts().subscribe(response => {
      console.log(response);
      this.products = response;
    })
  }
  sendRemoveProduct(item : http.RootObject) {
    this.httpService.deleteProduct(item).subscribe(response => {
      console.log(response);
    })
    this.ngOnInit();
  }
}
