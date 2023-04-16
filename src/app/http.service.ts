import { Injectable } from '@angular/core';
import { HttpClient, HttpHeaders, HttpResponse } from '@angular/common/http';
import { Observable } from 'rxjs';
import { environment } from '../environments/environment'

//   export interface Meta {
//       __typename: string;
//       query: string;
//       stackId: number;
//       stackType: string;
//       title: string;
//       categoryTitle: string;
//       totalItemCount: number;
//       layoutEnum: string;
//       adsBeacon?: any;
//       viewAllParams?: any;
//       fulfillmentIntent?: any;
//   }

//   export interface ImageInfo {
//       __typename: string;
//       thumbnailUrl: string;
//   }

//   export interface AvailabilityStatusV2 {
//       __typename: string;
//       value: string;
//       display: string;
//   }

//   export interface Badge {
//       __typename: string;
//       type: string;
//       dynamicDisplayName?: any;
//       beaconId?: any;
//   }

//   export interface Flag {
//       __typename: string;
//       id: string;
//       text: string;
//       key: string;
//   }

//   export interface Tag {
//       __typename: string;
//       id: string;
//       text: string;
//       key: string;
//   }

//   export interface UnifiedBadge {
//       __typename: string;
//       flags: Flag[];
//       tags: Tag[];
//   }

//   export interface PriceDisplayCodes {
//       __typename: string;
//       unitOfMeasure: string;
//       priceDisplayCondition?: any;
//       pricePerUnitUom: string;
//       unitPriceDisplayCondition?: any;
//       finalCostByWeight?: any;
//   }

//   export interface CurrentPrice {
//       __typename: string;
//       price: number;
//       priceString: string;
//   }

//   export interface ListPrice {
//       __typename: string;
//       price: number;
//       priceString: string;
//   }

//   export interface PriceInfo {
//       __typename: string;
//       priceRange?: any;
//       priceDisplayCodes: PriceDisplayCodes;
//       currentPrice: CurrentPrice;
//       wasPrice?: any;
//       unitPrice?: any;
//       listPrice: ListPrice;
//       subscriptionPrice?: any;
//       shipPrice?: any;
//   }

//   export interface PreOrder {
//       __typename: string;
//       isPreOrder: boolean;
//       preOrderMessage?: any;
//       preOrderStreetDateMessage?: any;
//   }

//   export interface Item {
//       __typename: string;
//       id: string;
//       usItemId: string;
//       name: string;
//       type: string;
//       showAtc: boolean;
//       showOptions: boolean;
//       checkStoreAvailabilityATC: boolean;
//       seeShippingEligibility: boolean;
//       annualEvent: boolean;
//       externalInfo?: any;
//       imageInfo: ImageInfo;
//       averageRating: number;
//       numberOfReviews: number;
//       salesUnitType: string;
//       esrb?: any;
//       mediaRunningTime?: any;
//       mediaRating?: any;
//       availabilityStatusV2: AvailabilityStatusV2;
//       mediaLanguage?: any;
//       orderLimit: number;
//       orderMinLimit: number;
//       weightIncrement: number;
//       badge: Badge[];
//       unifiedBadge: UnifiedBadge;
//       fulfillmentBadge: string;
//       fulfillmentSpeed: string[];
//       sellerId: string;
//       sellerName: string;
//       sponsoredProduct?: any;
//       offerId: string;
//       priceInfo: PriceInfo;
//       currencyCode: string;
//       fulfillmentType: string;
//       variantCriteria: any[];
//       preOrder: PreOrder;
//       fitmentLabel?: any;
//       productLocation?: any;
//       canonicalUrl: string;
//   }

//   export interface ItemStack {
//       __typename: string;
//       meta: Meta;
//       items: Item[];
//   }

//   export interface Location {
//       __typename: string;
//       addressId: string;
//   }

//   export interface PageMetadata {
//       __typename: string;
//       title: string;
//       noIndex: boolean;
//       location: Location;
//   }

//   export interface Spelling {
//       __typename: string;
//       correctedTerm?: any;
//   }

//   export interface PageProperties {
//       ptss: string;
//       ps: string;
//       prg: string;
//       spelling: string;
//       affinityOverride: string;
//       stores: string;
//       query: string;
//       cat_id: string;
//       sort: string;
//       displayGuidedNav: boolean;
//       page: number;
//   }

//   export interface PaginationV2 {
//       __typename: string;
//       maxPage: number;
//       pageProperties: PageProperties;
//       currentPage: number;
//       pap?: any;
//   }

//   export interface ErrorResponse {
//       __typename: string;
//       correlationId: string;
//       source: string;
//       errors: any[];
//   }

//   export interface RequestContext {
//       __typename: string;
//       searchMatchType: string;
//       shelfDisplayName: string;
//   }

//   export interface SearchResult {
//       __typename: string;
//       itemStacks: ItemStack[];
//       pageMetadata: PageMetadata;
//       spelling: Spelling;
//       paginationV2: PaginationV2;
//       errorResponse: ErrorResponse;
//       requestContext: RequestContext;
//       modules?: any;
//   }

//   export interface Search {
//       __typename: string;
//       query: string;
//       searchResult: SearchResult;
//   }

//   export interface Schedule {
//       __typename: string;
//       priority: number;
//   }

//   export interface MatchedTrigger {
//       __typename: string;
//       pageId: string;
//       zone: string;
//       inheritable: boolean;
//   }

//   export interface ViewConfig {
//       __typename: string;
//       title: string;
//       image: string;
//       displayName: string;
//       description: string;
//       url: string;
//       playStoreLink?: any;
//   }

//   export interface Image {
//       __typename: string;
//       src: string;
//   }

//   export interface PillsV2 {
//       __typename: string;
//       title: string;
//       catID?: any;
//       url: string;
//       suggestionType?: any;
//       catPathName?: any;
//       image: Image;
//   }

//   export interface ClickThrough {
//       __typename: string;
//       type: string;
//       value: string;
//   }

//   export interface Destination {
//       __typename: string;
//       title: string;
//       clickThrough: ClickThrough;
//   }

//   export interface Image2 {
//       __typename: string;
//       src: string;
//       alt: string;
//   }

//   export interface Campaigns {
//       __typename: string;
//       bannerBackgroundColor: string;
//       destination: Destination;
//       heading?: any;
//       subHeading?: any;
//       image: Image2;
//   }

//   export interface RawConfig {
//       moduleLocation: string;
//       enableLazyLoad: string;
//   }

//   export interface Value {
//       __typename: string;
//       id: string;
//       name: string;
//       layout?: any;
//       paramType?: any;
//       type: string;
//       url: string;
//       itemCount?: number;
//       expandOnLoad: boolean;
//       isSelected?: boolean;
//       min?: any;
//       max?: any;
//       selectedMin?: any;
//       selectedMax?: any;
//       unboundedMax?: any;
//       displayMultiLevelCategory?: any;
//       catPathName: string;
//       description: string;
//   }

//   export interface FacetsV1 {
//       __typename: string;
//       id?: any;
//       name: string;
//       layout: string;
//       paramType: string;
//       type: string;
//       url: string;
//       itemCount?: any;
//       expandOnLoad?: boolean;
//       isSelected?: any;
//       min?: number;
//       max?: number;
//       selectedMin?: any;
//       selectedMax?: any;
//       unboundedMax?: boolean;
//       displayMultiLevelCategory?: boolean;
//       catPathName?: any;
//       description?: any;
//       values: Value[];
//   }

//   export interface Configs {
//       __typename: string;
//       moduleType: string;
//       viewConfig: ViewConfig;
//       moduleSource: string;
//       pillsV2: PillsV2[];
//       campaigns: Campaigns;
//       ad?: any;
//       enableLazyLoad: string;
//       rawConfig: RawConfig;
//       fitments?: any;
//       facetsV1: FacetsV1[];
//   }

//   export interface Module {
//       __typename: string;
//       name: string;
//       version: number;
//       type: string;
//       moduleId: string;
//       schedule: Schedule;
//       matchedTrigger: MatchedTrigger;
//       configs: Configs;
//   }

//   export interface Location2 {
//       __typename: string;
//       stateOrProvinceCode: string;
//       postalCode: string;
//       storeId: string;
//   }

//   export interface Brand {
//       extractedValue: string;
//       pcs_brand: string[];
//       score: number;
//   }

//   export interface ProductType {
//       name: string;
//       score: number;
//       source: string;
//   }

//   export interface FeLog {
//       dept: string;
//       g: string;
//       s: string;
//       trf: string;
//   }

//   export interface AnalyticsLog {
//       fe_log: FeLog;
//   }

//   export interface SearchNormalize {
//       verticalId: string;
//       normalized_query: string;
//       original_query: string;
//       rewritten_query: string;
//       specificity: string;
//       top_query_cat_path: string;
//       top_query_cat_path_name: string;
//       brand: Brand;
//       product_type: ProductType[];
//       analytics_log: AnalyticsLog;
//   }

//   export interface PageContext {
//       searchNormalize: SearchNormalize;
//   }

//   export interface PageMetadata2 {
//       __typename: string;
//       location: Location2;
//       pageContext: PageContext;
//   }

//   export interface Options {
//       refId: string;
//       tempoLabel: string;
//       displayInTempo?: boolean;
//   }

//   export interface Mapping {
//       type: string;
//       options: Options;
//   }

//   export interface Content2 {
//       type: string;
//       id: string;
//       mapping: Mapping[];
//   }

//   export interface Content {
//       type: string;
//       flow: string;
//       content: Content2;
//   }

//   export interface Definition {
//       type: string;
//       flow: string;
//       content: Content[];
//   }

//   export interface Channel {
//       id: string;
//       status: string;
//   }

//   export interface Layout2 {
//       type: string;
//       name: string;
//       definition: Definition;
//       channel: Channel;
//       version: string;
//       status: string;
//   }

//   export interface Layout {
//       __typename: string;
//       id: string;
//       layout: Layout2;
//   }

//   export interface ContentLayout {
//       __typename: string;
//       modules: Module[];
//       pageMetadata: PageMetadata2;
//       layouts: Layout[];
//   }

//   export interface Data {
//       search: Search;
//       contentLayout: ContentLayout;
//   }

//   export interface RootObject {
//       data: Data;
//   }



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

    private bothURL = 'http://' + environment.serverURL + ':9000/bothStores'
    private signupUserURL = 'http://' + environment.serverURL + ':9000/api/user/signup'
    private loginUserURL = 'http://' + environment.serverURL + ':9000/api/user/signin'
    private authToken = '';
    private httpOptions = {
      observe: 'response' as 'response',
      headers: new HttpHeaders({
        'Content-Type':  'application/json',
        'Access-Control-Allow-Origin': '*',
      }),
      withCredentials: true
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
  }
