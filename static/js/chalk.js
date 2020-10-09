(function (root, factory) {
    if (typeof define === 'function' && define.amd) {
        // AMD. Register as an anonymous module.
        define(['exports', 'echarts'], factory);
    } else if (typeof exports === 'object' && typeof exports.nodeName !== 'string') {
        // CommonJS
        factory(exports, require('echarts'));
    } else {
        // Browser globals
        factory({}, root.echarts);
    }
}(this, function (exports, echarts) {
    var log = function (msg) {
        if (typeof console !== 'undefined') {
            console && console.error && console.error(msg);
        }
    };
    if (!echarts) {
        log('ECharts is not Loaded');
        return;
    }
    echarts.registerTheme('chalk', {
        "color": [
            "#f43ead",
            "#2396ff",
            "#7c43fc",
            "#ffda61"
        ],
        "backgroundColor": "rgba(41,52,65,0)",
        "textStyle": {},
        "title": {
            "textStyle": {
                "color": "#ffffff"
            },
            "subtextStyle": {
                "color": "#719bf2"
            }
        },
        "line": {
            "itemStyle": {
                "normal": {
                    "borderWidth": "4"
                }
            },
            "lineStyle": {
                "normal": {
                    "width": "3"
                }
            },
            "symbolSize": "0",
            "symbol": "circle",
            "smooth": true
        },
        "radar": {
            "itemStyle": {
                "normal": {
                    "borderWidth": "4"
                }
            },
            "lineStyle": {
                "normal": {
                    "width": "3"
                }
            },
            "symbolSize": "0",
            "symbol": "circle",
            "smooth": true
        },
        "bar": {
            "itemStyle": {
                "normal": {
                    "barBorderWidth": "1",
                    "barBorderColor": "#00ff28"
                },
                "emphasis": {
                    "barBorderWidth": "1",
                    "barBorderColor": "#00ff28"
                }
            }
        },
        "pie": {
            "itemStyle": {
                "normal": {
                    "borderWidth": "1",
                    "borderColor": "#00ff28"
                },
                "emphasis": {
                    "borderWidth": "1",
                    "borderColor": "#00ff28"
                }
            }
        },
        "scatter": {
            "itemStyle": {
                "normal": {
                    "borderWidth": "1",
                    "borderColor": "#00ff28"
                },
                "emphasis": {
                    "borderWidth": "1",
                    "borderColor": "#00ff28"
                }
            }
        },
        "boxplot": {
            "itemStyle": {
                "normal": {
                    "borderWidth": "1",
                    "borderColor": "#00ff28"
                },
                "emphasis": {
                    "borderWidth": "1",
                    "borderColor": "#00ff28"
                }
            }
        },
        "parallel": {
            "itemStyle": {
                "normal": {
                    "borderWidth": "1",
                    "borderColor": "#00ff28"
                },
                "emphasis": {
                    "borderWidth": "1",
                    "borderColor": "#00ff28"
                }
            }
        },
        "sankey": {
            "itemStyle": {
                "normal": {
                    "borderWidth": "1",
                    "borderColor": "#00ff28"
                },
                "emphasis": {
                    "borderWidth": "1",
                    "borderColor": "#00ff28"
                }
            }
        },
        "funnel": {
            "itemStyle": {
                "normal": {
                    "borderWidth": "1",
                    "borderColor": "#00ff28"
                },
                "emphasis": {
                    "borderWidth": "1",
                    "borderColor": "#00ff28"
                }
            }
        },
        "gauge": {
            "itemStyle": {
                "normal": {
                    "borderWidth": "1",
                    "borderColor": "#00ff28"
                },
                "emphasis": {
                    "borderWidth": "1",
                    "borderColor": "#00ff28"
                }
            }
        },
        "candlestick": {
            "itemStyle": {
                "normal": {
                    "color": "#fc97af",
                    "color0": "transparent",
                    "borderColor": "#fc97af",
                    "borderColor0": "#87f7cf",
                    "borderWidth": "2"
                }
            }
        },
        "graph": {
            "itemStyle": {
                "normal": {
                    "borderWidth": "1",
                    "borderColor": "#00ff28"
                }
            },
            "lineStyle": {
                "normal": {
                    "width": "1",
                    "color": "#ffffff"
                }
            },
            "symbolSize": "0",
            "symbol": "circle",
            "smooth": true,
            "color": [
                "#f43ead",
                "#2396ff",
                "#7c43fc",
                "#ffda61"
            ],
            "label": {
                "normal": {
                    "textStyle": {
                        "color": "#719bf2"
                    }
                }
            }
        },
        "map": {
            "itemStyle": {
                "normal": {
                    "areaColor": "#f3f3f3",
                    "borderColor": "#999999",
                    "borderWidth": 0.5
                },
                "emphasis": {
                    "areaColor": "rgba(255,178,72,1)",
                    "borderColor": "#eb8146",
                    "borderWidth": 1
                }
            },
            "label": {
                "normal": {
                    "textStyle": {
                        "color": "#893448"
                    }
                },
                "emphasis": {
                    "textStyle": {
                        "color": "rgb(137,52,72)"
                    }
                }
            }
        },
        "geo": {
            "itemStyle": {
                "normal": {
                    "areaColor": "#f3f3f3",
                    "borderColor": "#999999",
                    "borderWidth": 0.5
                },
                "emphasis": {
                    "areaColor": "rgba(255,178,72,1)",
                    "borderColor": "#eb8146",
                    "borderWidth": 1
                }
            },
            "label": {
                "normal": {
                    "textStyle": {
                        "color": "#893448"
                    }
                },
                "emphasis": {
                    "textStyle": {
                        "color": "rgb(137,52,72)"
                    }
                }
            }
        },
        "categoryAxis": {
            "axisLine": {
                "show": true,
                "lineStyle": {
                    "color": "rgba(102,102,102,0.34)"
                }
            },
            "axisTick": {
                "show": false,
                "lineStyle": {
                    "color": "#333"
                }
            },
            "axisLabel": {
                "show": true,
                "textStyle": {
                    "color": "#719bf2"
                }
            },
            "splitLine": {
                "show": true,
                "lineStyle": {
                    "color": [
                        "rgba(230,230,230,0.03)"
                    ]
                }
            },
            "splitArea": {
                "show": false,
                "areaStyle": {
                    "color": [
                        "rgba(250,250,250,0.05)",
                        "rgba(200,200,200,0.02)"
                    ]
                }
            }
        },
        "valueAxis": {
            "axisLine": {
                "show": true,
                "lineStyle": {
                    "color": "rgba(102,102,102,0.34)"
                }
            },
            "axisTick": {
                "show": false,
                "lineStyle": {
                    "color": "#333"
                }
            },
            "axisLabel": {
                "show": true,
                "textStyle": {
                    "color": "#719bf2"
                }
            },
            "splitLine": {
                "show": true,
                "lineStyle": {
                    "color": [
                        "rgba(230,230,230,0.03)"
                    ]
                }
            },
            "splitArea": {
                "show": false,
                "areaStyle": {
                    "color": [
                        "rgba(250,250,250,0.05)",
                        "rgba(200,200,200,0.02)"
                    ]
                }
            }
        },
        "logAxis": {
            "axisLine": {
                "show": true,
                "lineStyle": {
                    "color": "rgba(102,102,102,0.34)"
                }
            },
            "axisTick": {
                "show": false,
                "lineStyle": {
                    "color": "#333"
                }
            },
            "axisLabel": {
                "show": true,
                "textStyle": {
                    "color": "#719bf2"
                }
            },
            "splitLine": {
                "show": true,
                "lineStyle": {
                    "color": [
                        "rgba(230,230,230,0.03)"
                    ]
                }
            },
            "splitArea": {
                "show": false,
                "areaStyle": {
                    "color": [
                        "rgba(250,250,250,0.05)",
                        "rgba(200,200,200,0.02)"
                    ]
                }
            }
        },
        "timeAxis": {
            "axisLine": {
                "show": true,
                "lineStyle": {
                    "color": "rgba(102,102,102,0.34)"
                }
            },
            "axisTick": {
                "show": false,
                "lineStyle": {
                    "color": "#333"
                }
            },
            "axisLabel": {
                "show": true,
                "textStyle": {
                    "color": "#719bf2"
                }
            },
            "splitLine": {
                "show": true,
                "lineStyle": {
                    "color": [
                        "rgba(230,230,230,0.03)"
                    ]
                }
            },
            "splitArea": {
                "show": false,
                "areaStyle": {
                    "color": [
                        "rgba(250,250,250,0.05)",
                        "rgba(200,200,200,0.02)"
                    ]
                }
            }
        },
        "toolbox": {
            "iconStyle": {
                "normal": {
                    "borderColor": "#999999"
                },
                "emphasis": {
                    "borderColor": "#666666"
                }
            }
        },
        "legend": {
            "textStyle": {
                "color": "#999999"
            }
        },
        "tooltip": {
            "axisPointer": {
                "lineStyle": {
                    "color": "#cccccc",
                    "width": 1
                },
                "crossStyle": {
                    "color": "#cccccc",
                    "width": 1
                }
            }
        },
        "timeline": {
            "lineStyle": {
                "color": "#87f7cf",
                "width": 1
            },
            "itemStyle": {
                "normal": {
                    "color": "#00ffa5",
                    "borderWidth": 1
                },
                "emphasis": {
                    "color": "#f7f494"
                }
            },
            "controlStyle": {
                "normal": {
                    "color": "#87f7cf",
                    "borderColor": "#87f7cf",
                    "borderWidth": 0.5
                },
                "emphasis": {
                    "color": "#87f7cf",
                    "borderColor": "#87f7cf",
                    "borderWidth": 0.5
                }
            },
            "checkpointStyle": {
                "color": "#fc97af",
                "borderColor": "rgba(252,151,175,0.3)"
            },
            "label": {
                "normal": {
                    "textStyle": {
                        "color": "#87f7cf"
                    }
                },
                "emphasis": {
                    "textStyle": {
                        "color": "#87f7cf"
                    }
                }
            }
        },
        "visualMap": {
            "color": [
                "#f43ead",
                "#2396ff",
                "#7c43fc",
                "#ffda61"
            ]
        },
        "dataZoom": {
            "backgroundColor": "rgba(255,255,255,0)",
            "dataBackgroundColor": "rgba(114,204,255,1)",
            "fillerColor": "rgba(114,204,255,0.2)",
            "handleColor": "#72ccff",
            "handleSize": "100%",
            "textStyle": {
                "color": "#333333"
            }
        },
        "markPoint": {
            "label": {
                "normal": {
                    "textStyle": {
                        "color": "#719bf2"
                    }
                },
                "emphasis": {
                    "textStyle": {
                        "color": "#719bf2"
                    }
                }
            }
        }
    });
}));
