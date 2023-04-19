import { Component, ViewChild } from '@angular/core';
import { Router } from '@angular/router';
import { AppComponent } from '../app.component'
import { HttpService } from '../http.service';
import * as http from '../http.service';
import {MatSort, Sort} from '@angular/material/sort';
import {MatTableDataSource} from '@angular/material/table';
import { CookieService } from 'ngx-cookie-service';

@Component({
  selector: 'app-home',
  templateUrl: './home.component.html',
  styleUrls: ['./home.component.css']
})
export class HomeComponent {
  title = 'Home Page';
  items : http.ItemList;
  dataSource : MatTableDataSource<http.RootObject>;
  search : string;
  @ViewChild(MatSort) sort = new MatSort();
  displayedColumns : any[] = ['product_name','image_url','price','rating','seller_name', 'add_button'];
  showFiller = false;
  cookie_id = '';

  constructor(private router: Router, private httpService: HttpService, private cookie: CookieService) {}

  save() {
    // if the search term is empty, return -1
    if (this.search == "") {
      console.log(-1);
      return;
    }
    
    console.log(this.search);
    this.httpService.getAllItems(this.search).subscribe(response => {
      console.log(response);
      this.items = response;
      this.dataSource = new MatTableDataSource(this.items);
      this.dataSource.sort = this.sort;
      console.log(this.dataSource);
    })
  }

  sendSaveProduct(item : http.RootObject) {
    this.httpService.saveProduct(item).subscribe(response => {
      console.log(response);
      console.log(response.status);
    })
  }
  // routing function to take user to pageName
  goToPage(pageName:string):void {
    this.router.navigate([`${pageName}`]);
  }

  signOut() {
    // delete user cookie
    this.cookie.delete('token');
    window.location.reload();
  }

  ngAfterContentInit() {
    // get current cookie
    this.cookie_id = this.cookie.get('token');   
    console.log(this.cookie_id);
  }
}