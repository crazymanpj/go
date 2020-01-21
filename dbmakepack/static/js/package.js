/**
 * Created by kingsoft on 2018/1/16.
 */
var app = angular.module('package', []);
app.controller('conditionController', function ($scope, $http) {
    function init() {
        //1、获取产品列表
        //2、获取tryno列表
        //3、获取合作商列表

        $scope.models = ['非静默', '静默'];
        $scope.types = ['exe', 'dll'];

        $http({
            method: 'GET',
            url: urls.getproductlist
        }).success(function (data) {
            if (data.errorcode == 0) {
                $scope.products = data.productlist;
                $scope.product = $scope.products[0];
                $scope.panelType = allProducts[$scope.product];
                initDuba(false);
            }
        });
        // 30s更新一次历史列表
        setInterval(function () {
            //alert(1)
            refreshHistory();
        }, 1000 * 60);
    }

    init();
    $scope.product_change = function (x) {
        $scope.panelType = allProducts[x];
        initDuba(false);
        // if (x == "duba") {
        //     initDuba(false);
        // }
    };
    function initDuba(delay) {
        // 初始化其他字段
        $("#js-tid1").val('');
        $("#js-tid2").val('');
        $("#js-tod1").val('');
        $("#js-tod2").val('');
        $("#specialfile").val('');
        $("#localname").val('');
        $scope.packettype = "exe";
        $scope.packetmodel = "非静默";
        $scope.fixuplive = '否';
        $scope.islokmp = '是';

        // item 初始化
        initItemName();

        // 初始化installxml
        $http({
            method: 'GET',
            url: urls.getinstallxml + "?product=" + $scope.product
        }).success(function (data) {
            if (data.errorcode == 0) {
                $scope.installxmls = data.installxml;
                $scope.installxml = $scope.installxmls[0];
            } else {
                //请求错误
            }
        });

        // 初始化packetxml
        $http({
            method: 'GET',
            url: urls.getpacketxml + "?product=" + $scope.product
        }).success(function (data) {
            if (data.errorcode == 0) {
                $scope.packetxmls = data.packetxml;
                $scope.packetxml = $scope.packetxmls[0];
            } else {
                //请求错误
            }
        });

        // tryno初始化
        $http({
            method: 'GET',
            url: urls.gettyrnolist + "?product=" + $scope.product
        }).success(function (data) {
            if (data.errorcode == 0) {
                $scope.trynos = data.trynolist;
                $scope.tryno = $scope.trynos[0] + "";
            } else {
                //请求错误
            }
        }).error(function (data) {
            //访问超时
        });
        if (delay) {
            setTimeout(function () {
                refreshHistory();
            }, 1000 * 10);
        } else {
            refreshHistory();
        }

    }

    function refreshHistory() {
        // 历史记录
        console.log("刷新列表...");
        $http({
            method: 'GET',
            url: urls.gethistroypackages + "?product=" + $scope.product
        }).success(function (data) {
            if (data.errorcode == 0) {
                //console.log(data);
                $scope.historypackages = data.data;
            }
        });
    }


    $scope.itemname_change = function (itemname) {
        if (itemname == "add new ...") {
            $("#newItemnameModal").modal({
                keyboard: false,
                show: true,
                backdrop: 'static'
            });
        }
    };
    $scope.dismissNewItem = function () {
        $("#newItemnameModal").modal('hide');
        initItemName();

    };

    function initItemName(newItem) {
        $http({
            method: 'GET',
            url: urls.getitemlist + "?product=" + $scope.product
        }).success(function (data) {
            if (data.errorcode == 0) {
                addItemnameOptions(data.partnerlist, newItem);
                $scope.itemnames = data.partnerlist;
            } else {
                //请求错误

            }
        }).error(function (data) {
            //访问超时
        });
    }

    function addItemnameOptions(data, newItem) {
        var content = "";
        if (typeof(newItem) == "undefined") {
            $.each(data, function (idx, val) {
                if (idx == 1) {
                    content += "<option selected='true'>" + val + "</option>"
                } else {
                    content += "<option>" + val + "</option>"
                }
            });
        } else {
            $.each(data, function (idx, val) {

                if (val == newItem) {
                    content += "<option selected='true'>" + newItem + "</option>"
                } else {
                    content += "<option>" + val + "</option>";
                }
            });


        }
        content += "<option onclick='showNewItemName()'>add new ...</option>";
        $("#js-itemname-select").html(content);
    }

    $scope.saveNewItem = function () {
        var newitem = $("#newItemName").val();
        $http({
            method: 'post',
            url: urls.addnewpartner,
            data: {
                parner: newitem,
                product: $scope.product
            }
        }).success(function (data) {
            if (data.errorcode == 0) {
                $("#newItemnameModal").modal('hide');
                initItemName(newitem);
            } else {
                $scope.message = messages.error + "-- 错误码" + data.errorcode;
                $("#showMessageModal").modal('show');
                disappear("showMessageModal", 2000);
            }
        });

        //此处要加逻辑
        //调用增加item的接口
        //再刷新item控件
    };

    $scope.typechange = function (packettype) {
      if (packettype == "dll") {
          $scope.models = ['静默'];
          $scope.packetmodel = '静默';
      }
        if (packettype == "exe") {
            $scope.models = ['非静默', '静默'];
            $scope.packetmodel = '非静默';
        }
    };

    $scope.dopack = function () {
        $("#btn-pack").attr("disabled", "true");
        setTimeout(function () {
            $("#btn-pack").removeAttr("disabled");
        }, 1000 * 10);
        // 构造参数
        var data = {
            product: $("#js-product-select option:checked").text() + "",
            // product: "duba",
            itemname: $("#js-itemname-select option:checked").text(),
            // isnewitem: ($scope.itemnames.indexOf($("#js-itemname-select option:checked").text()) == -1 ? 1 : 0) + "",
            isnewitem: "1",
            // tryno: $("#js-tryno-select option:checked").text(),
            tryno: $("#js-tryno-select").val(),
            packettype: $("#js-packettype-select option:checked").val() + "",
            packetmodel: ($("#js-packetmodel-select option:checked").val() == "静默" ? 1 : 0) + "",
            tid1: $("#js-tid1").val() + "",
            tid2: $("#js-tid2").val() + "",
            tod2: $("#js-tod2").val() + "",
            tod1: $("#js-tod1").val() + "",
            fixuplive: ($("#js-fixuplive-select option:checked").val() == "是" ? 1 : 0) + "",
            islokmp: 0 + "", //这个参数暂时不需要
            specialfile: $("#specialfile").val(),
            localname: $("#localname").val(),
            installxml: $("#js-installxml-select option:checked").text(),
            packetxml: $("#js-packetxml-select option:checked").text()
        };
        // 没有做参数为空的校验
        $http({
            method: 'post',
            url: urls.dopack,
            data: data
        }).success(function (data) {
            if (data.errorcode == 0) {
                console.log("打包接口调用成功");
                //alert("开始打包..");
                $scope.message = messages.dopack;
                $("#showMessageModal").modal('show');
                disappear("showMessageModal", 2000);
                initDuba(true);
            } else if (data.errorcode == -1) {
                //有任务正在进行中...
                $scope.message = messages.working;
                $("#showMessageModal").modal('show');
                disappear("showMessageModal", 2000);
            } else {
                //请求错误
                $scope.message = messages.error + "-- 错误码" + data.errorcode;
                $("#showMessageModal").modal('show');
                disappear("showMessageModal", 2000);
            }
        }).error(function (data) {
            $scope.message = "打包参数非法。";
            $("#showMessageModal").modal('show');
            disappear("showMessageModal", 2000);
        });
    };

    $scope.stoppack = function (id) {
        $http({
            method: 'POST',
            url: urls.stoppack,
            data: {
                Taskid: id
            }
        }).success(function (data) {
            if (data.errorcode == 0) {
                // alert("停止成功")
                $scope.message = messages.stop;
                $("#showMessageModal").modal('show');
                disappear("showMessageModal", 2000);
                refreshHistory();
            } else {
                //请求错误
                $scope.message = messages.error + "-- 错误码" + data.errorcode;
                $("#showMessageModal").modal('show');
                disappear("showMessageModal", 2000);
            }
            refreshHistory();
        }).error(function (data) {
            //访问超时
        });
    };

    $scope.query = function (id) {
        $http({
            method: 'get',
            url: urls.getiffinish + "?Taskid=" + id
        }).success(function (data) {
            if (data.errorcode == 0) {
                // result 表示打包是否完成 0=未完成 1=完成
                if (data.result == 99) {
                    // alert("打包中...");
                    $scope.message = messages.packing;
                    $("#showMessageModal").modal('show');
                    disappear("showMessageModal", 2000);
                }
                if (data.result == 1) {
                    // alert("打包完成...");
                    $scope.message = messages.packed;
                    $("#showMessageModal").modal('show');
                    disappear("showMessageModal", 2000);
                    refreshHistory();
                }

            } else {
                //请求错误
                $scope.message = messages.error + "-- 错误码" + data.errorcode;
                $("#showMessageModal").modal('show');
                disappear("showMessageModal", 2000);
            }
        }).error(function (data) {
            //访问超时
        });
    };
    // $scope.showNewItemName = function (itemname) {
    //     // var itemname = $("#js-itemname-select option:checked").text()
    //     if (itemname == "add new ...") {
    //         $("#newItemnameModal").modal({
    //             keyboard: false,
    //             show: true,
    //             backdrop: 'static'
    //         });
    //     } else {
    //         // 这里要用jquery的请求方式
    //         console.log(itemname);
    //         var data = {
    //             itemname: itemname
    //         };
    //         console.log(data);
    //
    //     }
    // };
    $scope.showList = function () {
        var itemname = $("#js-itemname-select option:checked").text();
        $http({
            url: urls.getautolist,
            data: {
                itemname: itemname,
                product:  $scope.product
            },
            method: "post"
        }).success(function (data) {
            if (data.errorcode == 0) {
                // 请求成功
                $scope.selectItem = itemname;
                $scope.selectCount = data.data.length;
                $scope.selectData = data.data;
                $("#showListModal").modal({
                    show: true
                });
            } else {
                $scope.message = messages.error + "-- 错误码" + data.errorcode;
                $("#showMessageModal").modal('show');
                disappear("showMessageModal", 2000);
            }
        });
    };
    $scope.apply = function (x) {
        console.log(x);
        // 拿到了打包条件，下一步填充
        $("#js-tryno-select").val(x.tryno);
        $scope.packettype = x.packagetype;
        $("#js-packettype-select").find("option:contains('" + x.packagetype + "')").attr("selected", true);
        $scope.models = ['非静默', '静默'];
        $scope.packetmodel = x.packagemodel == 0 ? "非静默" : "静默";
        $("#js-tid1").val(x.tid1);
        $("#js-tid2").val(x.tid2);
        $("#js-tod1").val(x.tod1);
        $("#js-tod2").val(x.tod2);
        $scope.installxml = x.installxml;
        $scope.packetxml = x.packetxml;
        $("#localname").val(x.localname);
        // 还有点问题
        disappear("showListModal");
    }

});

function showNewItemName() {
    //alert("1");
    var itemname = $("#js-itemname-select option:checked").text();
    if (itemname == "add new ...") {
        $("#newItemnameModal").modal({
            keyboard: false,
            show: true,
            backdrop: 'static'
        });
    } else {
        // // 这里要用jquery的请求方式
        // console.log(itemname);
        // var data = {
        //     itemname: itemname
        // };
        // console.log(data);
        // $.post(urls.getautolist, data, function (data) {
        //     console.log(data);
        // })
    }
}

function copyDataPath(a) {
    var dataPath = $(a).attr('data');
    $(a).val(dataPath);
    a.select();
    document.execCommand("Copy");
    $(a).val('点击复制');
    $("#copy-success").show();
    setTimeout(function() {
        $("#copy-success").hide();
    }, 2000);
}

function disappear(id, delay) {
    if (typeof (delay) == "undefined") {
        $("#"+id).modal('hide');
    } else {
        setTimeout(function () {
            $("#" + id).modal('hide');
        }, delay);
    }
}