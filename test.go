


<input type="text" value="0" ng-keydown="quantityInputKeyDown($event)" ng-model="ticketModel.quantity" ng-keyup="quantityInputKeyUp($event)" ng-disabled="busy" class="ng-pristine ng-untouched ng-valid ng-not-empty">

<input id="person_agree_terms" type="checkbox" value="agree" ng-model="conditions.agreeTerm" class="ng-valid ng-dirty ng-valid-parse ng-not-empty ng-touched">

<button class="btn btn-primary btn-lg ng-isolate-scope" style="text-overflow:ellipsis; overflow: hidden;" kk-busy-spinner="busy &amp;&amp; !isBookingSkippable" ng-disabled="busy || !couldNextStep()" ng-class="{'btn-primary': couldNextStep(), 'btn-disabled-alt': !couldNextStep()}" ng-click="challenge()">

	<!-- ngIf: isBookingSkippable -->
	<!-- ngIf: !busy && !isBookingSkippable --><span ng-if="!busy &amp;&amp; !isBookingSkippable" class="ng-binding ng-scope">下一步</span><!-- end ngIf: !busy && !isBookingSkippable -->
	<!-- ngIf: busy && !isBookingSkippable -->
</button>

<button type="button" class="btn btn-primary dropdown-toggle ng-binding btnPrimary" ng-class="{btnPrimary: !showDropdown, btnDefault: showDropdown, active: showDropdown}" ng-click="showDropdown = !showDropdown">
確認座位 <span class="badge ng-binding">1</span>
<span class="caret"></span>
</button>

<a href="javascript:void(0)" class="btn btn-primary ng-binding" ng-click="done()">完成選位</a>

<a class="btn btn-primary btn-lg ng-binding ng-isolate-scope" kk-busy-spinner="submitting" ng-disabled="submitting" ng-click="confirmOrder()">確認表單資料</a>
