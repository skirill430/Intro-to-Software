import { Component } from '@angular/core';
import { Router } from '@angular/router';

@Component({
  selector: 'app-cart',
  templateUrl: './cart.component.html',
  styleUrls: ['./cart.component.css']
})
export class CartComponent {
  constructor (private router : Router) {}

  // Routing function to take user to Home Page
  goToPage(pageName:string):void {
    this.router.navigate([`${pageName}`]);
  }
}
