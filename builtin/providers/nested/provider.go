package nested

import (
	"fmt"
	"time"

	"github.com/davecgh/go-spew/spew"
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/terraform"
)

func Provider() terraform.ResourceProvider {
	return &schema.Provider{
		ResourcesMap: map[string]*schema.Resource{
			"nested": nestedSchema(),
		},
	}
}

func nestedSchema() *schema.Resource {
	return &schema.Resource{
		Create: nestedRead,
		Read:   nestedRead,
		Update: nestedRead,
		Delete: nestedRead,

		Schema: map[string]*schema.Schema{
			"rule": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"comment": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"criteria": {
							Type:     schema.TypeSet,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"name": {
										Type:     schema.TypeString,
										Required: true,
									},
									"option": {
										Type:     schema.TypeSet,
										Optional: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"name": {
													Type:     schema.TypeString,
													Required: true,
												},
												"values": {
													Type:     schema.TypeSet,
													Elem:     &schema.Schema{Type: schema.TypeString},
													Optional: true,
												},
												"value": {
													Type:     schema.TypeString,
													Optional: true,
												},
												"flag": {
													Type:     schema.TypeBool,
													Optional: true,
												},
											},
										},
									},
								},
							},
						},
						"behavior": {
							Type:     schema.TypeSet,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"name": {
										Type:     schema.TypeString,
										Required: true,
									},
									"option": {
										Type:     schema.TypeSet,
										Optional: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"name": {
													Type:     schema.TypeString,
													Required: true,
												},
												"values": {
													Type:     schema.TypeSet,
													Elem:     &schema.Schema{Type: schema.TypeString},
													Optional: true,
												},
												"value": {
													Type:     schema.TypeString,
													Optional: true,
												},
												"flag": {
													Type:     schema.TypeBool,
													Optional: true,
												},
											},
										},
									},
								},
							},
						},
						"rule": &schema.Schema{
							Type:     schema.TypeSet,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"name": {
										Type:     schema.TypeString,
										Required: true,
									},
									"comment": {
										Type:     schema.TypeString,
										Required: true,
									},
									"criteria": {
										Type:     schema.TypeSet,
										Optional: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"name": {
													Type:     schema.TypeString,
													Required: true,
												},
												"option": {
													Type:     schema.TypeSet,
													Optional: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"name": {
																Type:     schema.TypeString,
																Required: true,
															},
															"values": {
																Type:     schema.TypeSet,
																Elem:     &schema.Schema{Type: schema.TypeString},
																Optional: true,
															},
															"value": {
																Type:     schema.TypeString,
																Optional: true,
															},
															"flag": {
																Type:     schema.TypeBool,
																Optional: true,
															},
														},
													},
												},
											},
										},
									},
									"behavior": {
										Type:     schema.TypeSet,
										Optional: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"name": {
													Type:     schema.TypeString,
													Required: true,
												},
												"option": {
													Type:     schema.TypeSet,
													Optional: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"name": {
																Type:     schema.TypeString,
																Required: true,
															},
															"values": {
																Type:     schema.TypeSet,
																Elem:     &schema.Schema{Type: schema.TypeString},
																Optional: true,
															},
															"value": {
																Type:     schema.TypeString,
																Optional: true,
															},
															"flag": {
																Type:     schema.TypeBool,
																Optional: true,
															},
														},
													},
												},
											},
										},
									},
									"rule": &schema.Schema{
										Type:     schema.TypeSet,
										Optional: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"name": {
													Type:     schema.TypeString,
													Required: true,
												},
												"comment": {
													Type:     schema.TypeString,
													Required: true,
												},
												"criteria": {
													Type:     schema.TypeSet,
													Optional: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"name": {
																Type:     schema.TypeString,
																Required: true,
															},
															"option": {
																Type:     schema.TypeSet,
																Optional: true,
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{
																		"name": {
																			Type:     schema.TypeString,
																			Required: true,
																		},
																		"values": {
																			Type:     schema.TypeSet,
																			Elem:     &schema.Schema{Type: schema.TypeString},
																			Optional: true,
																		},
																		"value": {
																			Type:     schema.TypeString,
																			Optional: true,
																		},
																		"flag": {
																			Type:     schema.TypeBool,
																			Optional: true,
																		},
																	},
																},
															},
														},
													},
												},
												"behavior": {
													Type:     schema.TypeSet,
													Optional: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"name": {
																Type:     schema.TypeString,
																Required: true,
															},
															"option": {
																Type:     schema.TypeSet,
																Optional: true,
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{
																		"name": {
																			Type:     schema.TypeString,
																			Required: true,
																		},
																		"values": {
																			Type:     schema.TypeSet,
																			Elem:     &schema.Schema{Type: schema.TypeString},
																			Optional: true,
																		},
																		"value": {
																			Type:     schema.TypeString,
																			Optional: true,
																		},
																		"flag": {
																			Type:     schema.TypeBool,
																			Optional: true,
																		},
																	},
																},
															},
														},
													},
												},
												"rule": &schema.Schema{
													Type:     schema.TypeSet,
													Optional: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"name": {
																Type:     schema.TypeString,
																Required: true,
															},
															"comment": {
																Type:     schema.TypeString,
																Required: true,
															},
															"criteria": {
																Type:     schema.TypeSet,
																Optional: true,
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{
																		"name": {
																			Type:     schema.TypeString,
																			Required: true,
																		},
																		"option": {
																			Type:     schema.TypeSet,
																			Optional: true,
																			Elem: &schema.Resource{
																				Schema: map[string]*schema.Schema{
																					"name": {
																						Type:     schema.TypeString,
																						Required: true,
																					},
																					"values": {
																						Type:     schema.TypeSet,
																						Elem:     &schema.Schema{Type: schema.TypeString},
																						Optional: true,
																					},
																					"value": {
																						Type:     schema.TypeString,
																						Optional: true,
																					},
																					"flag": {
																						Type:     schema.TypeBool,
																						Optional: true,
																					},
																				},
																			},
																		},
																	},
																},
															},
															"behavior": {
																Type:     schema.TypeSet,
																Optional: true,
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{
																		"name": {
																			Type:     schema.TypeString,
																			Required: true,
																		},
																		"option": {
																			Type:     schema.TypeSet,
																			Optional: true,
																			Elem: &schema.Resource{
																				Schema: map[string]*schema.Schema{
																					"name": {
																						Type:     schema.TypeString,
																						Required: true,
																					},
																					"values": {
																						Type:     schema.TypeSet,
																						Elem:     &schema.Schema{Type: schema.TypeString},
																						Optional: true,
																					},
																					"value": {
																						Type:     schema.TypeString,
																						Optional: true,
																					},
																					"flag": {
																						Type:     schema.TypeBool,
																						Optional: true,
																					},
																				},
																			},
																		},
																	},
																},
															},
															"rule": &schema.Schema{
																Type:     schema.TypeSet,
																Optional: true,
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{
																		"name": {
																			Type:     schema.TypeString,
																			Required: true,
																		},
																		"comment": {
																			Type:     schema.TypeString,
																			Required: true,
																		},
																		"criteria": {
																			Type:     schema.TypeSet,
																			Optional: true,
																			Elem: &schema.Resource{
																				Schema: map[string]*schema.Schema{
																					"name": {
																						Type:     schema.TypeString,
																						Required: true,
																					},
																					"option": {
																						Type:     schema.TypeSet,
																						Optional: true,
																						Elem: &schema.Resource{
																							Schema: map[string]*schema.Schema{
																								"name": {
																									Type:     schema.TypeString,
																									Required: true,
																								},
																								"values": {
																									Type:     schema.TypeSet,
																									Elem:     &schema.Schema{Type: schema.TypeString},
																									Optional: true,
																								},
																								"value": {
																									Type:     schema.TypeString,
																									Optional: true,
																								},
																								"flag": {
																									Type:     schema.TypeBool,
																									Optional: true,
																								},
																							},
																						},
																					},
																				},
																			},
																		},
																		"behavior": {
																			Type:     schema.TypeSet,
																			Optional: true,
																			Elem: &schema.Resource{
																				Schema: map[string]*schema.Schema{
																					"name": {
																						Type:     schema.TypeString,
																						Required: true,
																					},
																					"option": {
																						Type:     schema.TypeSet,
																						Optional: true,
																						Elem: &schema.Resource{
																							Schema: map[string]*schema.Schema{
																								"name": {
																									Type:     schema.TypeString,
																									Required: true,
																								},
																								"values": {
																									Type:     schema.TypeSet,
																									Elem:     &schema.Schema{Type: schema.TypeString},
																									Optional: true,
																								},
																								"value": {
																									Type:     schema.TypeString,
																									Optional: true,
																								},
																								"flag": {
																									Type:     schema.TypeBool,
																									Optional: true,
																								},
																							},
																						},
																					},
																				},
																			},
																		},
																		"rule": &schema.Schema{
																			Type:     schema.TypeSet,
																			Optional: true,
																			Elem: &schema.Resource{
																				Schema: map[string]*schema.Schema{
																					"name": {
																						Type:     schema.TypeString,
																						Required: true,
																					},
																					"comment": {
																						Type:     schema.TypeString,
																						Required: true,
																					},
																					"criteria": {
																						Type:     schema.TypeSet,
																						Optional: true,
																						Elem: &schema.Resource{
																							Schema: map[string]*schema.Schema{
																								"name": {
																									Type:     schema.TypeString,
																									Required: true,
																								},
																								"option": {
																									Type:     schema.TypeSet,
																									Optional: true,
																									Elem: &schema.Resource{
																										Schema: map[string]*schema.Schema{
																											"name": {
																												Type:     schema.TypeString,
																												Required: true,
																											},
																											"values": {
																												Type:     schema.TypeSet,
																												Elem:     &schema.Schema{Type: schema.TypeString},
																												Optional: true,
																											},
																											"value": {
																												Type:     schema.TypeString,
																												Optional: true,
																											},
																											"flag": {
																												Type:     schema.TypeBool,
																												Optional: true,
																											},
																										},
																									},
																								},
																							},
																						},
																					},
																					"behavior": {
																						Type:     schema.TypeSet,
																						Optional: true,
																						Elem: &schema.Resource{
																							Schema: map[string]*schema.Schema{
																								"name": {
																									Type:     schema.TypeString,
																									Required: true,
																								},
																								"option": {
																									Type:     schema.TypeSet,
																									Optional: true,
																									Elem: &schema.Resource{
																										Schema: map[string]*schema.Schema{
																											"name": {
																												Type:     schema.TypeString,
																												Required: true,
																											},
																											"values": {
																												Type:     schema.TypeSet,
																												Elem:     &schema.Schema{Type: schema.TypeString},
																												Optional: true,
																											},
																											"value": {
																												Type:     schema.TypeString,
																												Optional: true,
																											},
																											"flag": {
																												Type:     schema.TypeBool,
																												Optional: true,
																											},
																										},
																									},
																								},
																							},
																						},
																					},
																					"rule": &schema.Schema{
																						Type:     schema.TypeSet,
																						Optional: true,
																						Elem: &schema.Resource{
																							Schema: map[string]*schema.Schema{
																								"name": {
																									Type:     schema.TypeString,
																									Required: true,
																								},
																								"comment": {
																									Type:     schema.TypeString,
																									Required: true,
																								},
																								"criteria": {
																									Type:     schema.TypeSet,
																									Optional: true,
																									Elem: &schema.Resource{
																										Schema: map[string]*schema.Schema{
																											"name": {
																												Type:     schema.TypeString,
																												Required: true,
																											},
																											"option": {
																												Type:     schema.TypeSet,
																												Optional: true,
																												Elem: &schema.Resource{
																													Schema: map[string]*schema.Schema{
																														"name": {
																															Type:     schema.TypeString,
																															Required: true,
																														},
																														"values": {
																															Type:     schema.TypeSet,
																															Elem:     &schema.Schema{Type: schema.TypeString},
																															Optional: true,
																														},
																														"value": {
																															Type:     schema.TypeString,
																															Optional: true,
																														},
																														"flag": {
																															Type:     schema.TypeBool,
																															Optional: true,
																														},
																													},
																												},
																											},
																										},
																									},
																								},
																								"behavior": {
																									Type:     schema.TypeSet,
																									Optional: true,
																									Elem: &schema.Resource{
																										Schema: map[string]*schema.Schema{
																											"name": {
																												Type:     schema.TypeString,
																												Required: true,
																											},
																											"option": {
																												Type:     schema.TypeSet,
																												Optional: true,
																												Elem: &schema.Resource{
																													Schema: map[string]*schema.Schema{
																														"name": {
																															Type:     schema.TypeString,
																															Required: true,
																														},
																														"values": {
																															Type:     schema.TypeSet,
																															Elem:     &schema.Schema{Type: schema.TypeString},
																															Optional: true,
																														},
																														"value": {
																															Type:     schema.TypeString,
																															Optional: true,
																														},
																														"flag": {
																															Type:     schema.TypeBool,
																															Optional: true,
																														},
																													},
																												},
																											},
																										},
																									},
																								},
																								"rule": &schema.Schema{
																									Type:     schema.TypeSet,
																									Optional: true,
																									Elem: &schema.Resource{
																										Schema: map[string]*schema.Schema{
																											"name": {
																												Type:     schema.TypeString,
																												Required: true,
																											},
																											"comment": {
																												Type:     schema.TypeString,
																												Required: true,
																											},
																											"criteria": {
																												Type:     schema.TypeSet,
																												Optional: true,
																												Elem: &schema.Resource{
																													Schema: map[string]*schema.Schema{
																														"name": {
																															Type:     schema.TypeString,
																															Required: true,
																														},
																														"option": {
																															Type:     schema.TypeSet,
																															Optional: true,
																															Elem: &schema.Resource{
																																Schema: map[string]*schema.Schema{
																																	"name": {
																																		Type:     schema.TypeString,
																																		Required: true,
																																	},
																																	"values": {
																																		Type:     schema.TypeSet,
																																		Elem:     &schema.Schema{Type: schema.TypeString},
																																		Optional: true,
																																	},
																																	"value": {
																																		Type:     schema.TypeString,
																																		Optional: true,
																																	},
																																	"flag": {
																																		Type:     schema.TypeBool,
																																		Optional: true,
																																	},
																																},
																															},
																														},
																													},
																												},
																											},
																											"behavior": {
																												Type:     schema.TypeSet,
																												Optional: true,
																												Elem: &schema.Resource{
																													Schema: map[string]*schema.Schema{
																														"name": {
																															Type:     schema.TypeString,
																															Required: true,
																														},
																														"option": {
																															Type:     schema.TypeSet,
																															Optional: true,
																															Elem: &schema.Resource{
																																Schema: map[string]*schema.Schema{
																																	"name": {
																																		Type:     schema.TypeString,
																																		Required: true,
																																	},
																																	"values": {
																																		Type:     schema.TypeSet,
																																		Elem:     &schema.Schema{Type: schema.TypeString},
																																		Optional: true,
																																	},
																																	"value": {
																																		Type:     schema.TypeString,
																																		Optional: true,
																																	},
																																	"flag": {
																																		Type:     schema.TypeBool,
																																		Optional: true,
																																	},
																																},
																															},
																														},
																													},
																												},
																											},
																											"rule": &schema.Schema{
																												Type:     schema.TypeSet,
																												Optional: true,
																												Elem: &schema.Resource{
																													Schema: map[string]*schema.Schema{
																														"name": {
																															Type:     schema.TypeString,
																															Required: true,
																														},
																														"comment": {
																															Type:     schema.TypeString,
																															Required: true,
																														},
																														"criteria": {
																															Type:     schema.TypeSet,
																															Optional: true,
																															Elem: &schema.Resource{
																																Schema: map[string]*schema.Schema{
																																	"name": {
																																		Type:     schema.TypeString,
																																		Required: true,
																																	},
																																	"option": {
																																		Type:     schema.TypeSet,
																																		Optional: true,
																																		Elem: &schema.Resource{
																																			Schema: map[string]*schema.Schema{
																																				"name": {
																																					Type:     schema.TypeString,
																																					Required: true,
																																				},
																																				"values": {
																																					Type:     schema.TypeSet,
																																					Elem:     &schema.Schema{Type: schema.TypeString},
																																					Optional: true,
																																				},
																																				"value": {
																																					Type:     schema.TypeString,
																																					Optional: true,
																																				},
																																				"flag": {
																																					Type:     schema.TypeBool,
																																					Optional: true,
																																				},
																																			},
																																		},
																																	},
																																},
																															},
																														},
																														"behavior": {
																															Type:     schema.TypeSet,
																															Optional: true,
																															Elem: &schema.Resource{
																																Schema: map[string]*schema.Schema{
																																	"name": {
																																		Type:     schema.TypeString,
																																		Required: true,
																																	},
																																	"option": {
																																		Type:     schema.TypeSet,
																																		Optional: true,
																																		Elem: &schema.Resource{
																																			Schema: map[string]*schema.Schema{
																																				"name": {
																																					Type:     schema.TypeString,
																																					Required: true,
																																				},
																																				"values": {
																																					Type:     schema.TypeSet,
																																					Elem:     &schema.Schema{Type: schema.TypeString},
																																					Optional: true,
																																				},
																																				"value": {
																																					Type:     schema.TypeString,
																																					Optional: true,
																																				},
																																				"flag": {
																																					Type:     schema.TypeBool,
																																					Optional: true,
																																				},
																																			},
																																		},
																																	},
																																},
																															},
																														},
																														"rule": &schema.Schema{
																															Type:     schema.TypeSet,
																															Optional: true,
																															Elem: &schema.Resource{
																																Schema: map[string]*schema.Schema{
																																	"name": {
																																		Type:     schema.TypeString,
																																		Required: true,
																																	},
																																	"comment": {
																																		Type:     schema.TypeString,
																																		Required: true,
																																	},
																																	"criteria": {
																																		Type:     schema.TypeSet,
																																		Optional: true,
																																		Elem: &schema.Resource{
																																			Schema: map[string]*schema.Schema{
																																				"name": {
																																					Type:     schema.TypeString,
																																					Required: true,
																																				},
																																				"option": {
																																					Type:     schema.TypeSet,
																																					Optional: true,
																																					Elem: &schema.Resource{
																																						Schema: map[string]*schema.Schema{
																																							"name": {
																																								Type:     schema.TypeString,
																																								Required: true,
																																							},
																																							"values": {
																																								Type:     schema.TypeSet,
																																								Elem:     &schema.Schema{Type: schema.TypeString},
																																								Optional: true,
																																							},
																																							"value": {
																																								Type:     schema.TypeString,
																																								Optional: true,
																																							},
																																							"flag": {
																																								Type:     schema.TypeBool,
																																								Optional: true,
																																							},
																																						},
																																					},
																																				},
																																			},
																																		},
																																	},
																																	"behavior": {
																																		Type:     schema.TypeSet,
																																		Optional: true,
																																		Elem: &schema.Resource{
																																			Schema: map[string]*schema.Schema{
																																				"name": {
																																					Type:     schema.TypeString,
																																					Required: true,
																																				},
																																				"option": {
																																					Type:     schema.TypeSet,
																																					Optional: true,
																																					Elem: &schema.Resource{
																																						Schema: map[string]*schema.Schema{
																																							"name": {
																																								Type:     schema.TypeString,
																																								Required: true,
																																							},
																																							"values": {
																																								Type:     schema.TypeSet,
																																								Elem:     &schema.Schema{Type: schema.TypeString},
																																								Optional: true,
																																							},
																																							"value": {
																																								Type:     schema.TypeString,
																																								Optional: true,
																																							},
																																							"flag": {
																																								Type:     schema.TypeBool,
																																								Optional: true,
																																							},
																																						},
																																					},
																																				},
																																			},
																																		},
																																	},
																																	"rule": &schema.Schema{
																																		Type:     schema.TypeSet,
																																		Optional: true,
																																		Elem: &schema.Resource{
																																			Schema: map[string]*schema.Schema{
																																				"name": {
																																					Type:     schema.TypeString,
																																					Required: true,
																																				},
																																				"comment": {
																																					Type:     schema.TypeString,
																																					Required: true,
																																				},
																																				"criteria": {
																																					Type:     schema.TypeSet,
																																					Optional: true,
																																					Elem: &schema.Resource{
																																						Schema: map[string]*schema.Schema{
																																							"name": {
																																								Type:     schema.TypeString,
																																								Required: true,
																																							},
																																							"option": {
																																								Type:     schema.TypeSet,
																																								Optional: true,
																																								Elem: &schema.Resource{
																																									Schema: map[string]*schema.Schema{
																																										"name": {
																																											Type:     schema.TypeString,
																																											Required: true,
																																										},
																																										"values": {
																																											Type:     schema.TypeSet,
																																											Elem:     &schema.Schema{Type: schema.TypeString},
																																											Optional: true,
																																										},
																																										"value": {
																																											Type:     schema.TypeString,
																																											Optional: true,
																																										},
																																										"flag": {
																																											Type:     schema.TypeBool,
																																											Optional: true,
																																										},
																																									},
																																								},
																																							},
																																						},
																																					},
																																				},
																																				"behavior": {
																																					Type:     schema.TypeSet,
																																					Optional: true,
																																					Elem: &schema.Resource{
																																						Schema: map[string]*schema.Schema{
																																							"name": {
																																								Type:     schema.TypeString,
																																								Required: true,
																																							},
																																							"option": {
																																								Type:     schema.TypeSet,
																																								Optional: true,
																																								Elem: &schema.Resource{
																																									Schema: map[string]*schema.Schema{
																																										"name": {
																																											Type:     schema.TypeString,
																																											Required: true,
																																										},
																																										"values": {
																																											Type:     schema.TypeSet,
																																											Elem:     &schema.Schema{Type: schema.TypeString},
																																											Optional: true,
																																										},
																																										"value": {
																																											Type:     schema.TypeString,
																																											Optional: true,
																																										},
																																										"flag": {
																																											Type:     schema.TypeBool,
																																											Optional: true,
																																										},
																																									},
																																								},
																																							},
																																						},
																																					},
																																				},
																																				"rule": &schema.Schema{
																																					Type:     schema.TypeSet,
																																					Optional: true,
																																					Elem: &schema.Resource{
																																						Schema: map[string]*schema.Schema{
																																							"name": {
																																								Type:     schema.TypeString,
																																								Required: true,
																																							},
																																							"comment": {
																																								Type:     schema.TypeString,
																																								Required: true,
																																							},
																																							"criteria": {
																																								Type:     schema.TypeSet,
																																								Optional: true,
																																								Elem: &schema.Resource{
																																									Schema: map[string]*schema.Schema{
																																										"name": {
																																											Type:     schema.TypeString,
																																											Required: true,
																																										},
																																										"option": {
																																											Type:     schema.TypeSet,
																																											Optional: true,
																																											Elem: &schema.Resource{
																																												Schema: map[string]*schema.Schema{
																																													"name": {
																																														Type:     schema.TypeString,
																																														Required: true,
																																													},
																																													"values": {
																																														Type:     schema.TypeSet,
																																														Elem:     &schema.Schema{Type: schema.TypeString},
																																														Optional: true,
																																													},
																																													"value": {
																																														Type:     schema.TypeString,
																																														Optional: true,
																																													},
																																													"flag": {
																																														Type:     schema.TypeBool,
																																														Optional: true,
																																													},
																																												},
																																											},
																																										},
																																									},
																																								},
																																							},
																																							"behavior": {
																																								Type:     schema.TypeSet,
																																								Optional: true,
																																								Elem: &schema.Resource{
																																									Schema: map[string]*schema.Schema{
																																										"name": {
																																											Type:     schema.TypeString,
																																											Required: true,
																																										},
																																										"option": {
																																											Type:     schema.TypeSet,
																																											Optional: true,
																																											Elem: &schema.Resource{
																																												Schema: map[string]*schema.Schema{
																																													"name": {
																																														Type:     schema.TypeString,
																																														Required: true,
																																													},
																																													"values": {
																																														Type:     schema.TypeSet,
																																														Elem:     &schema.Schema{Type: schema.TypeString},
																																														Optional: true,
																																													},
																																													"value": {
																																														Type:     schema.TypeString,
																																														Optional: true,
																																													},
																																													"flag": {
																																														Type:     schema.TypeBool,
																																														Optional: true,
																																													},
																																												},
																																											},
																																										},
																																									},
																																								},
																																							},
																																							"rule": &schema.Schema{
																																								Type:     schema.TypeSet,
																																								Optional: true,
																																								Elem: &schema.Resource{
																																									Schema: map[string]*schema.Schema{
																																										"name": {
																																											Type:     schema.TypeString,
																																											Required: true,
																																										},
																																										"comment": {
																																											Type:     schema.TypeString,
																																											Required: true,
																																										},
																																										"criteria": {
																																											Type:     schema.TypeSet,
																																											Optional: true,
																																											Elem: &schema.Resource{
																																												Schema: map[string]*schema.Schema{
																																													"name": {
																																														Type:     schema.TypeString,
																																														Required: true,
																																													},
																																													"option": {
																																														Type:     schema.TypeSet,
																																														Optional: true,
																																														Elem: &schema.Resource{
																																															Schema: map[string]*schema.Schema{
																																																"name": {
																																																	Type:     schema.TypeString,
																																																	Required: true,
																																																},
																																																"values": {
																																																	Type:     schema.TypeSet,
																																																	Elem:     &schema.Schema{Type: schema.TypeString},
																																																	Optional: true,
																																																},
																																																"value": {
																																																	Type:     schema.TypeString,
																																																	Optional: true,
																																																},
																																																"flag": {
																																																	Type:     schema.TypeBool,
																																																	Optional: true,
																																																},
																																															},
																																														},
																																													},
																																												},
																																											},
																																										},
																																										"behavior": {
																																											Type:     schema.TypeSet,
																																											Optional: true,
																																											Elem: &schema.Resource{
																																												Schema: map[string]*schema.Schema{
																																													"name": {
																																														Type:     schema.TypeString,
																																														Required: true,
																																													},
																																													"option": {
																																														Type:     schema.TypeSet,
																																														Optional: true,
																																														Elem: &schema.Resource{
																																															Schema: map[string]*schema.Schema{
																																																"name": {
																																																	Type:     schema.TypeString,
																																																	Required: true,
																																																},
																																																"values": {
																																																	Type:     schema.TypeSet,
																																																	Elem:     &schema.Schema{Type: schema.TypeString},
																																																	Optional: true,
																																																},
																																																"value": {
																																																	Type:     schema.TypeString,
																																																	Optional: true,
																																																},
																																																"flag": {
																																																	Type:     schema.TypeBool,
																																																	Optional: true,
																																																},
																																															},
																																														},
																																													},
																																												},
																																											},
																																										},
																																									},
																																								},
																																							},
																																						},
																																					},
																																				},
																																			},
																																		},
																																	},
																																},
																															},
																														},
																													},
																												},
																											},
																										},
																									},
																								},
																							},
																						},
																					},
																				},
																			},
																		},
																	},
																},
															},
														},
													},
												},
											},
										},
									},
								},
							},
						},
					},
				},
			},
		},
	}
}

func nestedRead(d *schema.ResourceData, meta interface{}) error {
	start := time.Now().UnixNano()
	s := d.Get("rule").(*schema.Set)
	dump := spew.Sdump(s)

	end := time.Now().UnixNano()

	delta := end - start

	return fmt.Errorf("Time to Get: %d\n\n%s\n\n", delta, dump)
}
