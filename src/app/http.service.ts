import { Injectable } from '@angular/core';
import { HttpClient, HttpHeaders, HttpResponse } from '@angular/common/http';
import { Observable } from 'rxjs';
import { environment } from '../environments/environment'

export interface RootObject {
	image_url: string;
	product_name: string;
	price: string;
	rating: string;
	seller_name: string;
}
  
export type ItemList = Array<RootObject>;

@Injectable({
  providedIn: 'root'
})

export class HttpService {

    private bothURL = 'https://' + environment.serverURL + ':9000/bothStores'
    private signupUserURL = 'https://' + environment.serverURL + ':9000/api/user/signup'
    private loginUserURL = 'https://' + environment.serverURL + ':9000/api/user/signin'
    private productURL = 'https://' + environment.serverURL + ':9000/api/products'
    private authToken = '';
    private httpOptions = {
      observe: 'response' as 'response',
      headers: new HttpHeaders({
        'Content-Type':  'application/json',
        'Access-Control-Allow-Origin': '*',
      }),
      withCredentials: true
    };
    private getOptions = {
      headers: new HttpHeaders({
        'Content-Type':  'application/json',
        'Access-Control-Allow-Origin': '*',
    }),
      withCredentials: true
    };
    private deleteOptions = {
      headers: new HttpHeaders({
        'Content-Type':  'application/json',
        'Access-Control-Allow-Origin': '*',
    }),
      withCredentials: true,
      body: {}
    };

  
    constructor(private http: HttpClient) {
    }
  
    getAllItems(query : String) : Observable<ItemList> {
      return this.http.post<ItemList>(this.bothURL, query);
    }
  
  //   sign up button call
    sendSignupInfo(username : String, password : String) : Observable<HttpResponse<any>>  {
      return this.http.post<HttpResponse<any>>(this.signupUserURL, { username: username, password: password }, this.httpOptions);
    }
  //   log in button call
    sendLoginInfo(username : String, password : String) : Observable<HttpResponse<any>>  {
      return this.http.post<HttpResponse<any>>(this.loginUserURL, { username: username, password: password }, this.httpOptions);
    }
    saveProduct(product: RootObject) {
      console.log(product)
      return this.http.post<HttpResponse<any>>(this.productURL, product, this.httpOptions);
    }
    deleteProduct(product: RootObject) {
      this.deleteOptions.body = product;
      return this.http.delete<HttpResponse<any>>(this.productURL, this.deleteOptions);
    }
    getProducts() : Observable<ItemList> {
      return this.http.get<ItemList>(this.productURL, this.getOptions);
    }
  }