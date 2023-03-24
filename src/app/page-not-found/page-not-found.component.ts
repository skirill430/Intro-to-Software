import { Component } from '@angular/core';
import { Router } from '@angular/router';

@Component({
  selector: 'app-page-not-found',
  templateUrl: './page-not-found.component.html',
  styleUrls: ['./page-not-found.component.css']
})
export class PageNotFoundComponent {
  title = 'Page Not Found';
  constructor(private router:Router){}

  // function to take user to home page
  goToPage(pageName:string):void {
    this.router.navigate([`${pageName}`]);
  }
}
